# generate.py is a script that generates the model and graphsql schema from the
# json data in the schema package, which follows the structure of
# schema/schema.json

# Module code
import inflection
import stringcase
from pathlib import Path
import subprocess
import os
from jsonschema import validate
import jsonschema
from re import A
import json


# put global constants here, to avoid magic strings
FILENAME = "generate.py"
GENERATED_MSG = "// ! This is a generated file, do not edit"
GO_DEFAULT_DATETIME_LAYOUT = "time.RFC3339Nano"
GO_EXT = ".go"
GRAPHQLS_EXT = ".graphqls"
GRAPHQLS_DIR = "graph/schema2"
MODEL_API = "api"
MODEL_PKG = "model2"
MODEL = "model"
MODEL_DESCRIPTION = "modelDescription"
MODEL_FIELDS = "modelFields"
MODEL_FIELD_DATETIME_LAYOUT = "datetimeLayout"
MODEL_FIELD_NAME = "name"
MODEL_FIELD_IDENTIFIER = "identifier"
MODEL_FIELD_GO_TYPE = "goType"
MODEL_FIELD_DESERIALIZER = "unmarshaller"
MODEL_TESTS = "modelTests"
MODEL_TEST_DESCRIPTION = "description"
MODEL_TEST_NAME = "name"
MODEL_TEST_JSON = "json"
SCHEMA_DIR = "schema"
SCHEMA_FILENAME = "schema.json"
ROOT_DIR = "/Users/richardvasquez/Developer/cql"

# get all of the relavent json files from the schema, all except the schema
schema_filenames = os.listdir()
schema_filenames.remove(SCHEMA_FILENAME)

# Try to follow the PEP 8 style guide:
# * https://www.python.org/dev/peps/pep-0008

# Try to keep functions alphabetical, with dependency exceptions.  In the case
# of a functional dependency, put the dependent function after the dependency
# function.


class Field:
    def __init__(self, field_dic: hash):
        self.datetime_layout = GO_DEFAULT_DATETIME_LAYOUT
        self.deserializer = ""
        self.identifier = field_dic[MODEL_FIELD_IDENTIFIER]
        self.go_type = field_dic[MODEL_FIELD_GO_TYPE]
        self.go_field_name = inflection.camelize(self.identifier)
        self.go_field_tag = inflection.camelize(
            self.identifier + "_json_tag", False)
        self.has_datetime_layout = False
        self.has_deserializer = False

        if MODEL_FIELD_DATETIME_LAYOUT in field_dic:
            self.has_datetime_layout = True
            self.datetime_layout = field_dic[MODEL_FIELD_DATETIME_LAYOUT]

        if MODEL_FIELD_DESERIALIZER in field_dic:
            self.has_deserializer = True
            self.deserializer = field_dic[MODEL_FIELD_DESERIALIZER]


class Scheme:
    def __init__(self, filename: str):
        self.api = ""
        self.description = ""
        self.go_comment = ""
        self.go_model_filename = ""
        self.go_model_name = ""
        self.go_model_variable_name = ""
        self.model = ""
        self.filename = filename
        self.fields = []

        chdir(SCHEMA_DIR)
        with open(self.filename) as json_file:
            data = json.load(json_file)
            is_valid = self.validate_json(data)
            if is_valid:
                self.api = data[MODEL_API]
                self.description = data[MODEL_DESCRIPTION]
                self.go_comment = _80_char_go_comment(self.description)
                self.go_model_name = stringcase.pascalcase(data[MODEL])
                self.go_model_variable_name = stringcase.camelcase(
                    self.go_model_name)
                self.model = data[MODEL]
                self.go_model_filename = Path(self.model).with_suffix(GO_EXT)
                for field_dic in data[MODEL_FIELDS]:
                    field = Field(field_dic)
                    self.fields.append(field)

    def schema(self):
        """This function loads the given schema available"""
        with open(SCHEMA_FILENAME, 'r') as file:
            return json.load(file)
        None

    def validate_json(self, data: hash):
        """REF: https://json-schema.org/ """
        try:
            validate(instance=data, schema=self.schema())
        except jsonschema.exceptions.ValidationError as err:
            print(err)
            return False
        return True


def _80_char_go_comment(comment: str):
    """
    _80_char_go_comment formats an 80 character go comment block from a
    string.
    """
    partitions = []
    line = "//"
    for word in comment.split(" "):
        if len(line + word) >= 80:
            line = line + "\n"
            partitions.append(line)
            line = "//"
        line = line + " " + word
    partitions.append(line)
    return " ".join(partitions)


def _80_char_graphqls_comment(comment: str):
    """
    _80_char_graphqls_comment formats an 80 character graphqls comment block
    from a string.
    """
    partitioned_comment = []
    line = ""
    for word in comment.split(" "):
        if len(line + word) >= 80:
            line = line + "\n"
            partitioned_comment.append(line)
            line = ""
        if len(line) > 0:
            line = line + " " + word
        else:
            line = word
    partitioned_comment.append(line)
    return "\n".join(['"""', "".join(partitioned_comment), '"""'])


def camelcase(string_to_camelcase: str):
    """
    camelcase will convert a string to camelcase (or Pascal case with a lower
    leading letter).  I.e. `ThisThing` becomes `thisThing`.  We use an
    underscore to ensure things like `ID` become `id` instead of `iD`
    """
    inflection.camelize(inflection.underscore(string_to_camelcase), False)


def chdir(dir: str):
    """
    chdir is used to change the working directory, first we chance to the root
    directory to ensure re-use, then we change the provided directory.
    """
    os.chdir(ROOT_DIR)
    os.chdir(dir)


def constantize_json_go_tags(scheme: Scheme):
    """
    constantize_json_go_tags will constantize all json fields for go to
    be used as the start of the function to avoid magic strings.
    """
    consts = []
    for field in scheme.fields:
        const = f'{field.go_field_tag} = "{field.identifier}"'
        consts.append(const)
    consts.sort()
    return f'const({";".join(consts)})'


def format_go(filename: str):
    """
    format_go will format a go file as well as implicitly include all required
    (missing) go imports.
    """
    subprocess.run(["go", "fmt"])
    subprocess.run(["goimports", "-w", filename])


def go_model_struct(scheme: Scheme):
    """
    go_model_struct will construct the the go code for a go struct, i.e.
    `type GoModel struct { Field1: string `json:field1` }`
    """
    fields = []
    for field in scheme.fields:
        literal = f"{field.go_field_name} {field.go_type}"
        literal += f'`json:"{field.identifier}"`'
        fields.append(literal)
    fields.sort()
    return f'type {scheme.go_model_name} struct {{{";".join(fields)}}}'


def go_model_struct_unmarshal_fn_sig(scheme: Scheme, field: Field):
    """
    go_model_struct_unmarshal_fn_sig provides the function signature for the
    go model's unmarshal method.
    """
    return f"{field.go_field_tag}, &{scheme.go_model_variable_name}.{field.go_field_name}"


def field_deserializer_defined(scheme: Scheme, field: Field):
    """
    field_deserialized_defined will return the go friendly unmarshal function
    if it is defined on the field json blob, otherwise it returns None.
    """
    sig = go_model_struct_unmarshal_fn_sig(scheme, field)
    if field.has_deserializer and field.deserializer == "UnmarshalFloatFromString":
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


def go_struct_access_var(scheme: Scheme, field: Field):
    """
    go_struct_access_var returns a string that represents a go struct var
    accessing one of it's fields with a dot, i.e. `myStruct.MyStructField`
    """
    return f"{scheme.go_model_variable_name}.{field.go_field_name}"


def go_time_unmarshal_block(scheme: Scheme, field: Field):
    """
    go_time_unmarshaller_block generates the code block for deeserializing data
    into a time.Time go object.
    """
    sig = []
    sig.append(field.datetime_layout)
    sig.append(field.go_field_tag)
    sig.append(f"&{go_struct_access_var(scheme, field)}")

    unmarshaller = []
    unmarshaller.append(f'err = data.UnmarshalTime({",".join(sig)})')
    unmarshaller.append(go_return_error())

    return ";".join(unmarshaller)


def field_deserializer_undefined(scheme: Scheme, field: Field):
    """
    field_deserializer_undefined gets the unmarshalX function names from the go
    types defined on a field blob
    """
    sig = go_model_struct_unmarshal_fn_sig(scheme, field)
    if field.go_type == "string":
        return f'\ndata.UnmarshalString({sig})'
    if field.go_type == "bool":
        return f'\ndata.UnmarshalBool({sig})'
    if field.go_type == "time.Time":
        return f"\n{go_time_unmarshal_block(scheme, field)}"
    return None


def go_model_deserializer_fn_calls(scheme: Scheme):
    """
    go_model_deserializer_fn_calls returns a string of all the unmarshalX
    function calls for a go model's UnmarshalJSON function.
    """
    deserializers = []
    for field in scheme.fields:
        defined = field_deserializer_defined(scheme, field)
        undefined = field_deserializer_undefined(scheme, field)
        deserializers.append(defined if defined != None else undefined)
    deserializers.sort()
    return ";".join(deserializers)


def go_model_struct_unmarshal_json_block(scheme: Scheme):
    """
    go_model_struct_json_unmarshaller create the UnmarshalJSON function for a
    go model's struct.
    """
    fn = []
    fn.append(constantize_json_go_tags(scheme))
    fn.append(go_model_json_deserializer_block())
    fn.append(go_model_deserializer_fn_calls(scheme))

    comment = f"UnmarshalJSON will deserialize bytes into a {scheme.go_model_name} model"
    go_comment = _80_char_go_comment(comment)

    reciever = f"\nfunc ({scheme.go_model_variable_name} *{scheme.go_model_name})"
    signature = f'UnmarshalJSON(d []byte) error {{{"".join(fn)}; return nil}}'
    function = reciever + signature
    return "".join([go_comment, function])


def create_go_model(scheme: Scheme):
    """generate the code for a go model"""
    chdir(MODEL_PKG)
    with open(scheme.go_model_filename, "w+") as goFile:
        goFile.write(" ".join(["package", MODEL_PKG]))
        goFile.write(f"\n{GENERATED_MSG}\n")
        goFile.write(f"\n{scheme.go_comment}")
        goFile.write(f"\n{go_model_struct(scheme)}")
        goFile.write(f"\n{go_model_struct_unmarshal_json_block(scheme)}")
        goFile.close()
        format_go(scheme.go_model_filename)


def graphqls_type(field: Field):
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
    return switch[field.go_type]


def json_go_type_dictionary(scheme: Scheme):
    """
    json_go_type_dictionary will return a dictionary where keys are the json
    formatted data and the values are the go struct field types.  i.e.
    {this_thing: string, that_think: float64}
    """
    dictionary = {}
    for field in scheme.fields:
        dictionary[field.identifier] = field.go_type
    return dictionary


# def create_go_model_test(scheme: Scheme, test):
#     name = test[MODEL_TEST_NAME]
#     description = test[MODEL_TEST_DESCRIPTION]
#     model_json = test[MODEL_TEST_JSON]

#     fn = []
#     fn.append("g := goblin.Goblin(t)")

#     goblin = []
#     goblin.append(f"response := []byte(`{model_json}`)")

#     expected = []
#     json_to_type = json_go_type_dictionary(data)
#     test = json.loads(model_json)
#     for field in data[MODEL_FIELDS]:
#         key = field[MODEL_FIELD_IDENTIFIER]
#         val = None
#         if json_to_type[key] == "string":
#             val = f'"{test[key]}"'
#         elif json_to_type[key] == "float64":
#             val = float(test[key])
#         elif json_to_type[key] == "bool":
#             val = "true" if test[key] else "false"

#         expected.append(f'{go_model_field_name(field)}: {val}')

#     expectedStr = ",".join(expected)
#     goblin.append(f"expected := {model_name}{{{expectedStr}}}")
#     goblin.append(f"actual := {model_name}{{}}")
#     goblin.append(
#         "if err := json.Unmarshal(response, &actual); err != nil {panic(err)}")

#     assertions = []
#     for field in data[MODEL_FIELDS]:
#         name = go_model_field_name(field)
#         assertions.append(f"g.Assert(actual.{name}).Eql(expected.{name})")

#     assertsionsStr = ";".join(assertions)
#     goblin.append(
#         f'g.It("should pass with all struct assertions", func(){{{assertsionsStr}}})')
#     goblinStr = ";".join(goblin)
#     fn.append(f'g.Describe("{description}", func() {{{goblinStr}}})')
#     fnStr = ";".join(fn)
#     return f"func Test{model_name}{name}(t *testing.T) {{{fnStr}}}"


# def create_go_model_tests(data):
#     if not MODEL_TESTS in data:
#         return

#     chdir(MODEL_PKG)
#     go_filename = Path(data[MODEL] + "_test").with_suffix(GO_EXT)

#     # open the go model file
#     go_file = open(go_filename, "w+")

#     # write data to the model file
#     go_file.write(" ".join(["package", MODEL_PKG]))
#     go_file.write(f"\n{GENERATED_MSG}\n")

#     for test in data[MODEL_TESTS]:
#         go_file.write(f"\n{create_go_model_test(data, test)}")

#     # close the go file
#     go_file.close()

#     # format the file and add necessary imports
#     format_go(go_filename)


# def create_graphqls_scheme(data):
#     chdir(GRAPHQLS_DIR)
#     go_filename = Path(data[MODEL]).with_suffix(GRAPHQLS_EXT)

#     # open the go model file
#     with open(go_filename, "w+") as go_file:
#         fields = []
#         for field in data[MODEL_FIELDS]:
#             fields.append(
#                 f"  {inflection.camelize(go_model_field_name(field), False)}: {graphqls_type(field)}")

#         fieldsStr = "\n".join(fields)
#         go_file.write(_80_char_graphqls_comment(data[MODEL_DESCRIPTION]))
#         go_file.write(
#             f"\ntype {go_model_struct_name(data)} {{\n{fieldsStr}\n}}\n")

#         # close the go file
#         go_file.close()

    # loop over the schemas to generate the model files
for filename in schema_filenames:
    # don't try to generate data from the generate.py file
    if filename == FILENAME:
        continue

    scheme = Scheme(filename)
    create_go_model(scheme)

# print(schema.api)

    # first let's read the filename into a dictionary
    # data = load_json(filename)
    # create_go_model(data)
    # create_go_model_tests(data)
    # create_graphqls_scheme(data)


# print(inflection.camelize(inflection.underscore("ThisField"), False))
