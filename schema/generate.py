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
import json
import yaml
from typing import List


# put global constants here, to avoid magic strings
FILENAME = "generate.py"
GRAPHQL_GENERATE_MSG = "# * This is a generated file, do not edit\n\n"
PROTOMODEL_GENERATED_MSG = "// * This is a generated file, do not edit"
PROTOMODEL_ACCESSORS_MSG = "// * This is a generated file, do not edit"
MODEL_GENERATED_MSG = "// * This file was initialized by schema/generate.py, but is open to extension"
ENDPOINT_FILENAME_PREFIX = "endpoint"
ENDPOINT_PATH = "path"
ENDPOINT_ENUM_ROOT = "enumRoot"
ENDPOINT_DESCRIPTION = "description"
ENDPOINT_QUERY_PARAMS = "queryParams"
ENDPOINT_PATH_COMMENT = "Get takes an endpoint const and endpoint arguments to parse the URL endpoint path."
GO_CLIENT = "client"
GO_DEFAULT_DATETIME_LAYOUT = "time.RFC3339Nano"
GO_ENDPOINT_ARGS = "EndpointArgs"
GO_ENDPOINT_TYPE = "Endpoint"
GO_EXT = ".go"
GRAPHQLS_EXT = ".graphqls"
GRAPHQLS_DIR = "graph/schema"
GQLGEN_DIR = "./"
GQLGEN_FILENAME = "gqlgen.yml"
GQLGEN_PROTOMODELS = "models"
MODEL_API = "api"
MODEL_ONLY = "modelOnly"
PROTOMODEL_PKG = "protomodel"
MODEL = "model"
MODEL_PKG = "model"
MODEL_DESCRIPTION = "modelDescription"
MODEL_FIELDS = "modelFields"
MODEL_FIELD_DATETIME_LAYOUT = "datetimeLayout"
MODEL_FIELD_NAME = "name"
MODEL_FIELD_IDENTIFIER = "identifier"
MODEL_FIELD_GO_TYPE = "goType"
MODEL_FIELD_DESERIALIZER = "unmarshaller"
MODEL_FIELD_DESCRIPTION = "description"
MODEL_TESTS = "modelTests"
MODEL_TEST_DESCRIPTION = "description"
MODEL_TEST_NAME = "name"
MODEL_TEST_JSON = "json"
SCHEMA_ENDPOINTS = "endpoints"
SCHEMA_DIR = os.getcwd()
SCHEMA_FILENAME = "schema.json"
SCHEMA_MODEL_DIR = f'{SCHEMA_DIR}/model'
ROOT_DIR = os.path.dirname(os.getcwd())
PROTO_ACCESSORS_FILENAME = "proto_accessors"

# get all of the relavent json files from the schema, all except the schema
schema_filenames = os.listdir(SCHEMA_MODEL_DIR)
schema_filenames.remove(SCHEMA_FILENAME)

# Try to follow the PEP 8 style guide:
# * https://www.python.org/dev/peps/pep-0008

# Try to keep functions alphabetical, with dependency exceptions.  In the case
# of a functional dependency, put the dependent function after the dependency
# function.


class PathPart:
    def __init__(self, name: str):
        # self.name = name.replace("{", "").replace("}", "")
        self.name = name
        self.go_var = inflection.camelize(self.name, False)
        self.go_arg = ""
        self.has_path_param = False

        if self.is_path_param():
            self.go_arg = f'*args["{self.name.replace("{", "").replace("}", "")}"].PathParam'
            self.has_path_param = True
        else:
            self.go_arg = f'"{name}"'

    def is_path_param(self):
        return "{" in self.name and "}" in self.name


class QueryParam:
    def __init__(self, query_param_dic):
        self.identifier = query_param_dic[MODEL_FIELD_IDENTIFIER]
        self.go_type = query_param_dic[MODEL_FIELD_GO_TYPE]
        self.go_proto_field_name = inflection.camelize(self.identifier)


class Endpoint:
    def __init__(self, api: str, endpoint_dic: hash):
        self.path = endpoint_dic[ENDPOINT_PATH]
        self.path_root = ""
        self.enum_root = endpoint_dic[ENDPOINT_ENUM_ROOT]
        self.go_const = f'{inflection.camelize(self.enum_root)}{GO_ENDPOINT_TYPE}'
        self.path_parts: List[PathPart] = []
        self.description = endpoint_dic[ENDPOINT_DESCRIPTION]
        self.has_query_params = False
        self.query_params: List[QueryParam] = []

        qp = f"{api}_{inflection.underscore(self.enum_root)}_options"
        self.graphqls_model_name = inflection.camelize(qp)
        self.graphql_query_param_filename = Path(qp).with_suffix(GRAPHQLS_EXT)

        if ENDPOINT_QUERY_PARAMS in endpoint_dic:
            self.has_query_params = True
            for dic in endpoint_dic[ENDPOINT_QUERY_PARAMS]:
                self.query_params.append(QueryParam(dic))

        first_part_set = False
        for part in self.path.split("/"):
            if part != '':
                if not first_part_set:
                    part = "/" + part
                    first_part_set = True
                self.path_parts.append(PathPart(part))

    def has_path_params(self):
        for part in self.path_parts:
            if part.has_path_param:
                return True
        return False


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
        self.description = ""
        self.struct_type = False

        if MODEL_FIELD_DATETIME_LAYOUT in field_dic:
            self.has_datetime_layout = True
            self.datetime_layout = field_dic[MODEL_FIELD_DATETIME_LAYOUT]

        if MODEL_FIELD_DESERIALIZER in field_dic:
            self.has_deserializer = True
            self.deserializer = field_dic[MODEL_FIELD_DESERIALIZER]

        if MODEL_FIELD_DESCRIPTION in field_dic:
            self.description = field_dic[MODEL_FIELD_DESCRIPTION]

        if self.non_struct_go_type(field_dic):
            self.go_proto_field_name = self.go_field_name
        else:
            self.go_proto_field_name = f"Proto{self.go_field_name}"
            self.struct_type = True

    @staticmethod
    def non_struct_go_type(field_dic: hash):
        """
        non_struct_go_type will return true if it's a basic go type, otherwise false
        """
        go_type = field_dic[MODEL_FIELD_GO_TYPE]
        if go_type == "string":
            return True
        if go_type == "bool":
            return True
        if go_type == "time.Time":
            return True
        if go_type == "int":
            return True
        if go_type == "[]string":
            return True
        if go_type == "float64":
            return True
        if "scalar" in go_type:
            return True
        return False


class Test:
    def __init__(self, test_dic: hash):
        self.name = test_dic[MODEL_TEST_NAME]
        self.description = test_dic[MODEL_TEST_DESCRIPTION]
        self.json = test_dic[MODEL_TEST_JSON]


class Scheme:
    def __init__(self, filename: str):
        self.api = ""
        self.description = ""
        self.go_comment = ""
        self.go_model_filename = ""
        self.go_model_name = ""
        self.model = ""
        self.modelOnly = False
        self.filename = filename
        self.go_test_filename = ""
        self.graphql_filename = ""
        self.fields = []
        self.tests = []
        self.endpoints: List[Endpoint] = []
        chdir(SCHEMA_MODEL_DIR)
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
                self.go_test_filename = Path(
                    f'{self.model}_test').with_suffix(GO_EXT)
                self.go_model_filename = Path(
                    self.model).with_suffix(GO_EXT)
                self.graphql_filename = Path(
                    self.model).with_suffix(GRAPHQLS_EXT)

                for field_dic in data[MODEL_FIELDS]:
                    field = Field(field_dic)
                    self.fields.append(field)

                if SCHEMA_ENDPOINTS in data:
                    for endpoint_dic in data[SCHEMA_ENDPOINTS]:
                        endpoint = Endpoint(self.api, endpoint_dic)
                        self.endpoints.append(endpoint)

                if MODEL_TESTS in data:
                    for test_dic in data[MODEL_TESTS]:
                        test = Test(test_dic)
                        self.tests.append(test)

                if MODEL_ONLY in data:
                    self.modelOnly = data[MODEL_ONLY]

    def protoapi_dir(self):
        return self.api

    @staticmethod
    def schema():
        """This function loads the given schema available"""
        with open(SCHEMA_FILENAME, 'r') as file:
            return json.load(file)

    def validate_json(self, data: hash):
        """REF: https://json-schema.org/ """
        try:
            validate(instance=data, schema=self.schema())
        except jsonschema.exceptions.ValidationError as err:
            raise Exception(err)
        return True


class EndpointSchemaAPISet:
    def __init__(self):
        self.dic = {}
        return

    def add(self, scheme: Scheme):
        if scheme.api not in self.dic:
            self.dic[scheme.api] = []
        for endpoint in scheme.endpoints:
            self.dic[scheme.api].append(endpoint)

    def sort(self):
        for _, endpoints in self.dic.items():
            endpoints.sort(key=lambda x: x.enum_root)

        # if enum_roots exist, sort them and then append endpoints
        # alphabetically
        # if len(enum_roots) > 0:
        #     enum_roots.sort()
        #     for enum_root in enum_roots:
        #         for endpoint in tmp_endpoints:
        #             if endpoint.enum_root == enum_root:
        #                 self.endpoints.append(endpoint)


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


def protomodel_package_name():
    return f"package {PROTOMODEL_PKG};"


def endpoint_package_name(api: str):
    return f"package {api};"


def model_package_name():
    return f"package {MODEL_PKG};"


def model_generated_message():
    return f"\n{MODEL_GENERATED_MSG}\n"


def protomodel_generated_message():
    return f"\n{PROTOMODEL_GENERATED_MSG}\n"


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


def generate_graphqls():
    chdir(GQLGEN_DIR)
    subprocess.run(["go", "run", "github.com/99designs/gqlgen", "generate"])


def go_protomodel_struct(scheme: Scheme):
    """
    go_protomodel_struct will construct the the go code for a go struct, i.e.
    `type GoModel struct { Field1: string `json:field1` }`
    """
    fields = []
    for field in scheme.fields:
        literal = ""
        if field.description != "":
            literal += f"\n// {field.description}\n"
        literal += f"{field.go_proto_field_name} {field.go_type}"
        literal += f'`json:"{field.identifier}"`'
        fields.append(literal)
    fields.sort()
    joined = "\n".join(fields)
    return f'type {scheme.go_model_name} struct {{{joined}}}'


def go_model_struct(scheme: Scheme):
    """
    go_model_struct will construct the the go code for a go struct, i.e.
    `type GoModel struct { protomodel.GoModel }`
    """
    return f'type {scheme.go_model_name} struct {{{PROTOMODEL_PKG}.{scheme.go_model_name}}}'


def go_protomodel_struct_unmarshal_fn_sig(scheme: Scheme, field: Field):
    """
    go_protomodel_struct_unmarshal_fn_sig provides the function signature for the
    go model's unmarshal method.
    """
    return f"{field.go_field_tag}, &{scheme.go_model_variable_name}.{field.go_proto_field_name}"


def field_deserializer_defined(scheme: Scheme, field: Field):
    """
    field_deserialized_defined will return the go friendly unmarshal function
    if it is defined on the field json blob, otherwise it returns None.
    """
    sig = go_protomodel_struct_unmarshal_fn_sig(scheme, field)
    if field.has_deserializer and field.deserializer == "UnmarshalFloatFromString":
        return f'\ndata.UnmarshalFloatFromString({sig})'
    return None


def go_return_error():
    return "if err != nil { return err }"


def go_panic_error():
    return "if err != nil { panic(err) }"


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
    return f"{scheme.go_model_variable_name}.{field.go_proto_field_name}"


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
    sig = go_protomodel_struct_unmarshal_fn_sig(scheme, field)
    if field.deserializer != "":
        return f'\ndata.{field.deserializer}({sig})'
    if field.go_type == "string":
        return f'\ndata.UnmarshalString({sig})'
    if field.go_type == "bool":
        return f'\ndata.UnmarshalBool({sig})'
    if field.go_type == "time.Time":
        return f"\n{go_time_unmarshal_block(scheme, field)}"
    if field.go_type == "int":
        return f"\ndata.UnmarshalInt({sig})"
    if field.go_type == "[]string":
        return f"\ndata.UnmarshalStringSlice({sig})"
    if field.go_type == "float64":
        return f"\ndata.UnmarshalFloat({sig})"
    if "scalar" in field.go_type:
        return f'\ndata.Unmarshal{field.go_type.replace("scalar.", "")}({sig})'

    # if the go type contains slice brackets like [], then we need to unmarshal
    # as as slice
    if "[]" in field.go_type:
        accessor = f'{scheme.go_model_variable_name}.{field.go_proto_field_name}'
        data_accessor = f'data.Value({field.go_field_tag})'
        typ = field.go_type.replace("[]", "").replace("*", "")
        marshal_logic = f'byt, _ := json.Marshal(item); obj := {typ}{{}};'
        unmarshal_logic = 'if err := json.Unmarshal(byt, &obj); err != nil { return err };'
        logic = f'{marshal_logic} {unmarshal_logic} {accessor} = append({accessor}, &obj)'
        loop = f'for _, item := range {data_accessor}.([]interface{{}}) {{{logic}}}'
        cond = f'\nif v := {data_accessor}; v != nil {{{loop}}}'
        return cond
    else:
        # if the type is undefined, then it's a struct reference to a model-only
        # json
        accessor = f'\n{scheme.go_model_variable_name}.{field.go_proto_field_name}'
        init = f'{accessor} = {field.go_type}{{}}'
        return f'{init}; if err := data.UnmarshalStruct({sig}); err != nil {{return err}}'
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


def go_protomodel_struct_unmarshal_json_block(scheme: Scheme):
    """
    go_protomodel_struct_json_unmarshaller create the UnmarshalJSON function for a
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


def create_go_protomodel(scheme: Scheme):
    """generate the code for a go model"""
    chdir(PROTOMODEL_PKG)
    with open(scheme.go_model_filename, "w+") as go_file:
        go_file.write(protomodel_package_name())
        go_file.write(protomodel_generated_message())
        go_file.write(f"\n{scheme.go_comment}")
        go_file.write(f"\n{go_protomodel_struct(scheme)}")
        go_file.write(f"\n{go_protomodel_struct_unmarshal_json_block(scheme)}")
        go_file.close()
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
        "int": "Int",
        "[]string": "[String]"
    }
    if field.go_type in switch:
        return switch[field.go_type]
    if "[]" in field.go_type:
        return f'[{field.go_type.replace("[]", "").replace("*", "")}]'
    return field.go_type.replace("scalar.", "")


def tag_to_go_type(scheme: Scheme):
    """
    tag_to_go_type will return a dictionary where keys are the json formatted
    data and the values are the go struct field types.  i.e.
    {this_thing: string, that_think: float64}
    """
    dictionary = {}
    for field in scheme.fields:
        dictionary[field.identifier] = field.go_type
    return dictionary


def goblin_assertions(scheme: Scheme):
    """
    goblin_assertions writes the assertion list for the model we're trying to
    test, comparing the expected value passed by the user in the json schema
    to the actual value after deserialization
    """
    assertions = []
    for field in scheme.fields:
        name = field.go_proto_field_name
        assertions.append(f"g.Assert(actual.{name}).Eql(expected.{name})")
    return ";".join(assertions)


def goblin_deserializer(scheme: Scheme):
    """
    goblin_deserializer will generate the code used to deserialize the test json
    into a model instance
    """
    deserializer = []
    deserializer.append("err := json.Unmarshal(response, &actual)")
    deserializer.append(go_panic_error())
    msg = "should pass with all struct assertions"
    return f'g.It("{msg}", func(){{{";".join(deserializer)}; {goblin_assertions(scheme)}}})'


def goblin_expected(scheme: Scheme, test: Test):
    """
    go_model_test_expected_instance will create the expected instance (of a
    go struct) for testing.  It's what we will compare the unmarshalled data to.
    """
    partitions = []
    tag_go_type = tag_to_go_type(scheme)
    test_json = json.loads(test.json)
    for field in scheme.fields:
        tag = field.identifier
        val = None
        if tag_go_type[tag] == "string":
            val = f'"{test_json[tag]}"'
        elif tag_go_type[tag] == "float64":
            val = float(test_json[tag])
        elif tag_go_type[tag] == "bool":
            val = "true" if test_json[tag] else "false"
        elif tag_go_type[tag] == "int":
            val = int(test_json[tag])
        partitions.append(f'{field.go_proto_field_name}: {val}')
    return f'{scheme.go_model_name}{{{",".join(partitions)}}}'


def goblin_test(scheme: Scheme, test: Test):
    """
    goblin_test generates the code to test the scheme-defined tests, comparing
    an expected struct to an actual struct per the schema.
    """
    goblin = []
    goblin.append(f"response := []byte(`{test.json}`)")
    goblin.append(f"expected := {goblin_expected(scheme, test)}")
    goblin.append(f"actual := {scheme.go_model_name}{{}}")
    goblin.append(goblin_deserializer(scheme))
    return ";".join(goblin)


def create_go_protomodel_test(scheme: Scheme, test: Test):
    """
    create_go_protomodel_test will generate a go file for a single goblin test using
    the expected data from the schema
    """
    fn = []
    fn.append("g := goblin.Goblin(t)")
    fn.append(
        f'g.Describe("{test.description}", func() {{{goblin_test(scheme, test)}}})')
    return f'func Test{scheme.go_model_name}{test.name}(t *testing.T) {{{";".join(fn)}}}'


def create_go_protomodel_tests(scheme: Scheme):
    """
    create_go_protomodel_tests will generate all the goblin tests in the tests array
    in the model schema.
    """
    if len(scheme.tests) == 0:
        return

    chdir(PROTOMODEL_PKG)
    with open(scheme.go_test_filename, "w+") as go_file:
     # write data to the model file
        go_file.write(protomodel_package_name())
        go_file.write(protomodel_generated_message())
        for test in scheme.tests:
            go_file.write(f"\n{create_go_protomodel_test(scheme, test)}")
        go_file.close()
    format_go(scheme.go_test_filename)


def create_graphqls_scheme(scheme: Scheme):
    """
    create_graphqls_scheme will create the scheme for the graphqls object that
    should be 1-1 with the corresponding protomodel.
    """
    chdir(GRAPHQLS_DIR)
    with open(scheme.graphql_filename, "w+") as graphql_file:
        fields = []
        for field in scheme.fields:
            fields.append(
                f"  {inflection.camelize(field.go_field_name, False)}: {graphqls_type(field)}")

        fieldsStr = "\n".join(fields)
        graphql_file.write(GRAPHQL_GENERATE_MSG)
        graphql_file.write(_80_char_graphqls_comment(scheme.description))
        graphql_file.write(
            f"\ntype {scheme.go_model_name} {{\n{fieldsStr}\n}}\n")

        graphql_file.close()


def create_graphqls_query_params(endpoint: Endpoint):
    if not endpoint.has_query_params:
        return
    chdir(GRAPHQLS_DIR)
    with open(endpoint.graphql_query_param_filename, "w+") as graphql_file:
        params = []
        for qp in endpoint.query_params:
            params.append(
                f"  {inflection.camelize(qp.go_proto_field_name, False)}: {graphqls_type(qp)}")

        paramsStr = "\n".join(params)
        graphql_file.write(GRAPHQL_GENERATE_MSG)
        graphql_file.write(
            f"input {endpoint.graphqls_model_name} {{\n{paramsStr}\n}}\n")

        graphql_file.close()


def gqlgen_protomodel(scheme: Scheme):
    """
    gqlgen_protomodels will return the list of protomodels for gqlgen
    """
    return {"model": f'cql/{MODEL_PKG}.{scheme.go_model_name}'}


def update_gqlgen(scheme: Scheme):
    """
    update_gqlgen will update the gqlgen.yml file with data from the current
    schema
    """
    chdir(GQLGEN_DIR)
    with open(GQLGEN_FILENAME) as gqlgen_file:
        configs = yaml.load(gqlgen_file, Loader=yaml.Loader)
        if GQLGEN_PROTOMODELS not in configs:
            configs[GQLGEN_PROTOMODELS] = {}
        configs[GQLGEN_PROTOMODELS][scheme.go_model_name] = gqlgen_protomodel(
            scheme)
        stream = open(GQLGEN_FILENAME, 'w')
        yaml.dump(configs, stream)


# def create_protomodel_fn(scheme: Scheme):
#     fn = ""
#     receiver = f"{scheme.go_model_variable_name} *{scheme.go_model_name}"
#     return_val = f"*protomodel.{scheme.go_model_name}"
#     fn += f'\nfunc ({receiver}) Protomodel() {return_val} {{&{return_val}{{}}}}'


def create_go_model(scheme: Scheme):
    """
    create_go_model will create the extendable go model
    """
    chdir(MODEL_PKG)
    if os.path.exists(scheme.go_model_filename):
        return

    with open(scheme.go_model_filename, "w+") as go_file:
        go_file.write(model_package_name())
        go_file.write(model_generated_message())
        go_file.write(f"\n{scheme.go_comment}")
        go_file.write(f"\n{go_model_struct(scheme)}")
        go_file.close()
        format_go(scheme.go_model_filename)


def endpoint_filename():
    return Path(ENDPOINT_FILENAME_PREFIX).with_suffix(GO_EXT)


def endpoint_type():
    return f"\ntype {GO_ENDPOINT_TYPE} int;"


def endpoint_consts(endpoints: List[Endpoint]):
    consts = []
    initializer_const = f'_ {GO_ENDPOINT_TYPE} = iota'
    consts.append(initializer_const)
    for endpoint in endpoints:
        consts.append(endpoint.go_const)
    return f'const({";".join(consts)});'


def endpoint_path_fn_name(endpoint: Endpoint):
    return f'{endpoint.enum_root}Path'


def endpoint_get_fn_map_wrapper(mappings):
    _map = f"map[{GO_ENDPOINT_TYPE}]func(args {GO_CLIENT}.{GO_ENDPOINT_ARGS}) string"
    joined = "\n".join(mappings)
    return f'{_map}{{\n{joined}\n}}'


def endpoint_get_fn_wrapper(mappings):
    comment = _80_char_go_comment(ENDPOINT_PATH_COMMENT)
    rec = f"\n{comment}\nfunc (endpoint {GO_ENDPOINT_TYPE})"
    sig = f"Path(args {GO_CLIENT}.{GO_ENDPOINT_ARGS}) string"
    return f'{rec} {sig} {{return {endpoint_get_fn_map_wrapper(mappings)}[endpoint](args)}};'


def endpoing_get_fn_join_paths(endpoint: Endpoint):
    paths = []
    for part in endpoint.path_parts:
        paths.append(part.go_arg)
    joined = ",".join(paths)
    return f"path.Join({joined})"


def endpoint_get_fn(endpoints: List[Endpoint]):
    mappings = []
    for endpoint in endpoints:
        mappings.append(
            f'{endpoint.go_const}: {endpoint_path_fn_name(endpoint)},')
    return endpoint_get_fn_wrapper(mappings)


def endpoint_path_fns(endpoints: List[Endpoint]):
    fns = []
    for endpoint in endpoints:
        var = "_"
        if endpoint.has_path_params():
            var = "args"

        sig = f'func {endpoint_path_fn_name(endpoint)}({var} {GO_CLIENT}.{GO_ENDPOINT_ARGS}) '
        inner_logic = ""
        if endpoint.has_query_params:
            sig += "(p string)"
            inner_logic = f'p = {endpoing_get_fn_join_paths(endpoint)};'
            inner_logic += "var sb strings.Builder;"
            inner_logic += "sb.WriteString(p);"
            inner_logic += "sb.WriteString(args.QueryPath().String());"
            inner_logic += "return sb.String();"
        else:
            sig += "string"
            inner_logic = f'return {endpoing_get_fn_join_paths(endpoint)}'

        logic = f'{{ \n {inner_logic}  }}'
        desc = _80_char_go_comment(endpoint.description)
        fns.append(f'\n{desc}\n {sig} {logic}')
    return "\n\n".join(fns)


def create_endpoints(api: str, endpoints: List[Endpoint]):
    chdir(api)
    filename = endpoint_filename()
    with open(filename, "w+") as go_file:
        go_file.write(endpoint_package_name(api))
        go_file.write(protomodel_generated_message())
        go_file.write(endpoint_type())
        go_file.write(endpoint_consts(endpoints))
        go_file.write(endpoint_path_fns(endpoints))
        go_file.write(endpoint_get_fn(endpoints))
        go_file.close()
        format_go(filename)


def proto_accessors_filename():
    return Path(PROTO_ACCESSORS_FILENAME).with_suffix(GO_EXT)


def write_proto_accessors(scheme: Scheme):
    accessors = []
    reciever = f"m *{scheme.go_model_name}"
    for field in scheme.fields:
        if field.struct_type:
            return_type = f"*{field.go_type}"
            comment = f"{field.go_field_name} converts the native protomodel {field.go_proto_field_name} to a local {field.go_type}.\n"
            fn = "\n"
            fn += _80_char_go_comment(comment)

            logic = ""
            if "[]" in return_type:
                return_type = f"(slice {field.go_type})"
                base_model = field.go_type.replace("[]", "").replace("*", "")
                append_slice = f"slice = append(slice, &{base_model}{{{base_model}: *v}})"
                logic += f"for _, v := range m.{field.go_proto_field_name} {{{append_slice}}};"
                logic += "return"
            else:
                logic = f"return &{field.go_type}{{\n{field.go_type}: m.{field.go_proto_field_name}}}"

            fn += f"func ({reciever}) {field.go_field_name}() {return_type} {{{logic}}};"
            accessors.append(fn)
    return "".join(accessors)


def create_proto_accessors(_schema: List[Scheme]):
    chdir(MODEL_PKG)
    filename = proto_accessors_filename()
    with open(filename, "w+") as go_file:
        go_file.write(model_package_name())
        go_file.write(protomodel_generated_message())
        for scheme in _schema:
            go_file.write(write_proto_accessors(scheme))
        go_file.close()
        format_go(filename)


def generate_models():
    endpoint_set = EndpointSchemaAPISet()
    _schema: List[Scheme] = []
    for filename in schema_filenames:
        scheme = Scheme(filename)
        if not scheme.modelOnly:
            endpoint_set.add(scheme)
        create_go_protomodel(scheme)
        create_go_protomodel_tests(scheme)
        create_go_model(scheme)
        create_graphqls_scheme(scheme)
        update_gqlgen(scheme)
        _schema.append(scheme)

    create_proto_accessors(_schema)

    endpoint_set.sort()
    for api, endpoints in endpoint_set.dic.items():
        create_endpoints(api, endpoints)
        for endpoint in endpoints:
            create_graphqls_query_params(endpoint)


generate_models()
generate_graphqls()
