# generate.py is a script that generates the model and graphsql schema from the
# json data in the schema package, which follows the structure of
# schema/schema.json

import json
from re import A
import jsonschema
from jsonschema import validate
import os
import subprocess
from pathlib import Path
import stringcase
import inflection

# put global constants here, to avoid magic strings
GENERATED_MSG = "// ! This is a generated file, do not edit"
GO_DEFAULT_DATETIME_LAYOUT = "time.RFC3339Nano"
GO_EXT = ".go"
GRAPHQLS_EXT = ".graphqls"
GRAPHQLS_DIR = "graph/schema2"
MODEL_PKG = "model2"
MODEL = "model"
MODEL_DESCRIPTION = "modelDescription"
MODEL_FIELDS = "modelFields"
MODEL_FIELD_DATETIME_LAYOUT = "datetimeLayout"
MODEL_FIELD_NAME = "name"
MODEL_FIELD_IDENTIFIER = "identifier"
MODEL_FIELD_TYPE = "type"
MODEL_FIELD_UNMARSHALLER = "unmarshaller"
MODEL_TESTS = "modelTests"
MODEL_TEST_DESCRIPTION = "description"
MODEL_TEST_NAME = "name"
MODEL_TEST_JSON = "json"
SCHEMA_DIR = "schema"
SCHEMA_FILENAME = "schema.json"
ROOT_DIR = "/Users/richardvasquez/Developer/cql"

# get all of the relavent json files from the schema, all except the schema
schema_filenames = os.listdir(SCHEMA_DIR)
schema_filenames.remove(SCHEMA_FILENAME)

# Try to follow the PEP 8 style guide:
# * https://www.python.org/dev/peps/pep-0008

# Try to keep functions alphabetical, with dependency exceptions.  In the case
# of a functional dependency, put the dependent function after the dependency
# function.


def _80_char_go_comment(str):
    """
    _80_char_go_comment formats an 80 character go comment block from a
    string.
    """
    comment = []
    line = f"//"
    for word in str.split(" "):
        if len(line + word) >= 80:
            line = line + "\n"
            comment.append(line)
            line = "//"
        line = line + " " + word
    comment.append(line)
    return " ".join(comment)


def _80_char_graphqls_comment(str):
    """
    _80_char_graphqls_comment formats an 80 character graphqls comment block
    from a string.
    """
    comment = []
    line = ""
    for word in str.split(" "):
        if len(line + word) >= 80:
            line = line + "\n"
            comment.append(line)
            line = ""
        if len(line) > 0:
            line = line + " " + word
        else:
            line = word
    comment.append(line)
    commentStr = "".join(comment)
    return "\n".join(['"""', commentStr, '"""'])


def camelcase(str):
    """
    camelcase will convert a string to camelcase (or Pascal case with a lower
    leading letter).  I.e. `ThisThing` becomes `thisThing`.  We use an
    underscore to ensure things like `ID` become `id` instead of `iD`
    """
    inflection.camelize(inflection.underscore(str), False)


def chdir(dir):
    """
    chdir is used to change the working directory, first we chance to the root
    directory to ensure re-use, then we change the provided directory.
    """
    os.chdir(ROOT_DIR)
    os.chdir(dir)


def constantize_json_go_tag(field):
    """
    constantize_json_go_tag returns a json name fully capitalized, i.e. a constant
    in our go repo: SOMETHING_LIKE_THIS.
    """
    return inflection.camelize("json_tag_"+field[MODEL_FIELD_IDENTIFIER], False)


def constantize_json_go_tags(data):
    """
    constantize_json_go_tags will constantize all json fields for go to
    be used as the start of the function to avoid magic strings.
    """
    consts = []
    for field in data[MODEL_FIELDS]:
        const = f'{constantize_json_go_tag(field)} = "{field[MODEL_FIELD_IDENTIFIER]}"'
        consts.append(const)
    consts.sort()
    return f'const({";".join(consts)})'


def format_go(filename):
    """
    format_go will format a go file as well as implicitly include all required
    (missing) go imports.
    """
    subprocess.run(["go", "fmt"])
    subprocess.run(["goimports", "-w", filename])


def go_datetime_layout(field):
    """
    go_datetime_layout will return the time.Time layout used to unmarshal
    json datetime data into a time.Time go object.  Check the global constant
    for the default.
    """
    if MODEL_FIELD_DATETIME_LAYOUT in field:
        return field[MODEL_FIELD_DATETIME_LAYOUT]
    return GO_DEFAULT_DATETIME_LAYOUT


def go_model_field_name(field):
    return inflection.camelize(field[MODEL_FIELD_IDENTIFIER])


def go_model_struct_name(data):
    return stringcase.pascalcase(data[MODEL])


def go_model_struct(data):
    """
    go_model_struct will construct the the go code for a go struct, i.e.
    `type GoModel struct { Field1: string `json:field1` }`
    """
    fields = []
    for field in data[MODEL_FIELDS]:
        literal = f"{go_model_field_name(field)} {field[MODEL_FIELD_TYPE]}"
        literal += f'`json:"{field[MODEL_FIELD_IDENTIFIER]}"`'
        fields.append(literal)
    fields.sort()
    return f'type {go_model_struct_name(data)} struct {{{";".join(fields)}}}'


def go_model_var_name(data):
    return stringcase.camelcase(go_model_struct_name(data))


def go_model_struct_unmarshal_fn_sig(data, field):
    """
    go_model_struct_unmarshal_fn_sig provides the function signature for the
    go model's unmarshal method.
    """
    go_field_name = go_model_field_name(field)
    go_var_name = go_model_var_name(data)
    return f"{constantize_json_go_tag(field)}, &{go_var_name}.{go_field_name}"


def field_deserializer_defined(data, field):
    """
    field_deserialized_defined will return the go friendly unmarshal function
    if it is defined on the field json blob, otherwise it returns None.
    """
    sig = go_model_struct_unmarshal_fn_sig(data, field)
    if MODEL_FIELD_UNMARSHALLER in field:
        deserializer = field[MODEL_FIELD_UNMARSHALLER]
        if deserializer == "UnmarshalFloatFromString":
            return f'\ndata.UnmarshalFloatFromString({sig})'
    return None


def go_return_error():
    return "if err != nil { return err }"


def go_model_json_deserializer_block():
    """
    go_deserializer_block returns the block of code responsible for creating
    the data bytes stream from a json transform, used to unmarshal data.
    """
    return f"\ndata, err := serial.NewJSONTransform(d); {go_return_error()}"


def go_struct_access_var(data, field):
    """
    go_struct_access_var returns a string that represents a go struct var
    accessing one of it's fields with a dot, i.e. `myStruct.MyStructField`
    """
    return f"{go_model_var_name(data)}.{go_model_field_name(field)}"


def go_time_unmarshal_block(data, field):
    """
    go_time_unmarshaller_block generates the code block for deeserializing data
    into a time.Time go object.
    """
    sig = []
    sig.append(go_datetime_layout(field))
    sig.append(constantize_json_go_tag(field))
    sig.append(f"&{go_struct_access_var(data, field)}")

    unmarshaller = []
    unmarshaller.append(f'err = data.UnmarshalTime({",".join(sig)})')
    unmarshaller.append(go_return_error())

    return ";".join(unmarshaller)


def field_deserializer_undefined(data, field):
    """
    field_deserializer_undefined gets the unmarshalX function names from the go
    types defined on a field blob
    """
    sig = go_model_struct_unmarshal_fn_sig(data, field)
    go_type = field[MODEL_FIELD_TYPE]
    if go_type == "string":
        return f'\ndata.UnmarshalString({sig})'
    if go_type == "bool":
        return f'\ndata.UnmarshalBool({sig})'
    if go_type == "time.Time":
        return f"\n{go_time_unmarshal_block(data, field)}"
    return None


def go_model_deserializer_fn_calls(data):
    """
    go_model_deserializer_fn_calls returns a string of all the unmarshalX
    function calls for a go model's UnmarshalJSON function.
    """
    deserializers = []
    for field in data[MODEL_FIELDS]:
        defined = field_deserializer_defined(data, field)
        undefined = field_deserializer_undefined(data, field)
        deserializers.append(defined if defined != None else undefined)
    deserializers.sort()
    return ";".join(deserializers)


def go_model_struct_unmarshal_json_block(data):
    """
    go_model_struct_json_unmarshaller create the UnmarshalJSON function for a
    go model's struct.
    """
    go_model = go_model_struct_name(data)
    go_var_name = go_model_var_name(data)

    fn = []
    fn.append(constantize_json_go_tags(data))
    fn.append(go_model_json_deserializer_block())
    fn.append(go_model_deserializer_fn_calls(data))

    comment = f"UnmarshalJSON will deserialize bytes into a {go_model} model"
    go_comment = _80_char_go_comment(comment)

    reciever = f"\nfunc ({go_var_name} *{go_model})"
    signature = f'UnmarshalJSON(d []byte) error {{{"".join(fn)}; return nil}}'
    function = reciever + signature
    return "".join([go_comment, function])


def graphqls_type(field):
    """
    graphqls_type converts the go type specified on a field blob to a
    graphqls type.
    """
    switch = {
        "string": "String",
        "float64": "Float",
        "bool": "Boolean",
        "time.Time": "Time",
    }
    return switch[field[MODEL_FIELD_TYPE]]


def json_go_type_dictionary(data):
    """
    json_go_type_dictionary will return a dictionary where keys are the json
    formatted data and the values are the go struct field types.  i.e.
    {this_thing: string, that_think: float64}
    """
    d = {}
    for field in data[MODEL_FIELDS]:
        d[field[MODEL_FIELD_IDENTIFIER]] = field[MODEL_FIELD_TYPE]
    return d


def load_json(filename):
    """
    load_json will load the json data using the json filename from the schema
    dir. It retuns None if no associated json file is found.
    """
    chdir(SCHEMA_DIR)
    with open(filename) as json_file:
        data = json.load(json_file)
        is_valid = validate_json(data)
        if is_valid:
            return data
    None


def schema():
    """This function loads the given schema available"""
    with open(SCHEMA_FILENAME, 'r') as file:
        return json.load(file)
    None


def validate_json(data):
    """REF: https://json-schema.org/ """
    try:
        validate(instance=data, schema=schema())
    except jsonschema.exceptions.ValidationError as err:
        print(err)
        return False
    return True


def create_go_model(data):
    """generate the code for a go model"""
    chdir(MODEL_PKG)
    model_name = data[MODEL]
    go_model_name = stringcase.pascalcase(model_name)
    model_desc = data[MODEL_DESCRIPTION]
    go_filename = Path(data[MODEL]).with_suffix(GO_EXT)

    # open the go model file
    go_file = open(go_filename, "w+")

    # write data to the model file
    go_file.write(" ".join(["package", MODEL_PKG]))
    go_file.write(f"\n{GENERATED_MSG}\n")
    go_file.write(f"\n{_80_char_go_comment(model_desc)}")
    go_file.write(f"\n{go_model_struct(data)}")
    go_file.write(f"\n{go_model_struct_unmarshal_json_block(data)}")

    # close the go file
    go_file.close()

    # format the file and add necessary imports
    format_go(go_filename)


def create_go_model_test(data, test):
    name = test[MODEL_TEST_NAME]
    description = test[MODEL_TEST_DESCRIPTION]
    model_name = go_model_struct_name(data)
    model_json = test[MODEL_TEST_JSON]

    fn = []
    fn.append("g := goblin.Goblin(t)")

    goblin = []
    goblin.append(f"response := []byte(`{model_json}`)")

    expected = []
    json_to_type = json_go_type_dictionary(data)
    test = json.loads(model_json)
    for field in data[MODEL_FIELDS]:
        key = field[MODEL_FIELD_IDENTIFIER]
        val = None
        if json_to_type[key] == "string":
            val = f'"{test[key]}"'
        elif json_to_type[key] == "float64":
            val = float(test[key])
        elif json_to_type[key] == "bool":
            val = "true" if test[key] else "false"

        expected.append(f'{go_model_field_name(field)}: {val}')

    expectedStr = ",".join(expected)
    goblin.append(f"expected := {model_name}{{{expectedStr}}}")
    goblin.append(f"actual := {model_name}{{}}")
    goblin.append(
        "if err := json.Unmarshal(response, &actual); err != nil {panic(err)}")

    assertions = []
    for field in data[MODEL_FIELDS]:
        name = go_model_field_name(field)
        assertions.append(f"g.Assert(actual.{name}).Eql(expected.{name})")

    assertsionsStr = ";".join(assertions)
    goblin.append(
        f'g.It("should pass with all struct assertions", func(){{{assertsionsStr}}})')
    goblinStr = ";".join(goblin)
    fn.append(f'g.Describe("{description}", func() {{{goblinStr}}})')
    fnStr = ";".join(fn)
    return f"func Test{model_name}{name}(t *testing.T) {{{fnStr}}}"


def create_go_model_tests(data):
    if not MODEL_TESTS in data:
        return

    chdir(MODEL_PKG)
    go_filename = Path(data[MODEL] + "_test").with_suffix(GO_EXT)

    # open the go model file
    go_file = open(go_filename, "w+")

    # write data to the model file
    go_file.write(" ".join(["package", MODEL_PKG]))
    go_file.write(f"\n{GENERATED_MSG}\n")

    for test in data[MODEL_TESTS]:
        go_file.write(f"\n{create_go_model_test(data, test)}")

    # close the go file
    go_file.close()

    # format the file and add necessary imports
    format_go(go_filename)


def create_graphqls_scheme(data):
    chdir(GRAPHQLS_DIR)
    go_filename = Path(data[MODEL]).with_suffix(GRAPHQLS_EXT)

    # open the go model file
    go_file = open(go_filename, "w+")

    fields = []
    for field in data[MODEL_FIELDS]:
        fields.append(
            f"  {inflection.camelize(go_model_field_name(field), False)}: {graphqls_type(field)}")

    fieldsStr = "\n".join(fields)
    go_file.write(_80_char_graphqls_comment(data[MODEL_DESCRIPTION]))
    go_file.write(
        f"\ntype {go_model_struct_name(data)} {{\n{fieldsStr}\n}}\n")

    # close the go file
    go_file.close()


# loop over the schemas to generate the model files
for filename in schema_filenames:
    # first let's read the filename into a dictionary
    print(filename)
    data = load_json(filename)
    create_go_model(data)
    create_go_model_tests(data)
    create_graphqls_scheme(data)


print(inflection.camelize(inflection.underscore("ThisField"), False))
