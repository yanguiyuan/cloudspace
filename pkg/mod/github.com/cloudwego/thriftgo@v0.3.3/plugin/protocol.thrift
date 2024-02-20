// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

namespace * plugin

include "../parser/AST.thrift"

// Request is the input of a plugin serialized in binary protocol.
struct Request {
    // The version of thiftgo.
    1: required string Version,

    // The parameters for generator from the command line, in the form of 'key=val' or 'val'.
    2: required list<string> GeneratorParameters,

    // The parameters for plugin from the command line, in the form of 'key=val' or 'val'.
    3: required list<string> PluginParameters,

    // The target language.
    4: required string Language,

    // The output path of generated files.
    5: required string OutputPath,

    // Recursively process the included files.
    6: required bool Recursive,

    // The abstract syntax trees of the parsed thrift IDL.
    7: required AST.Thrift AST,
}

struct Generated {
    // The content generated by the plugin.
    1: required string Content,

    // The target filename that `Content` is supposed to be written or inserted.
    // If not set, the content will be appended to the file specified by last `Generated`.
    // If set, the filename should be an absolute path.
    2: optional string Name,

    // If set, indicates the insertion point that the content is supposed to be inserted before.
    3: optional string InsertionPoint,
}

// Response is the output of a plugin serialized in binary protocol.
struct Response {
    // Error message.
    1: optional string Error,

    // Generated contents.
    2: optional list<Generated> Contents,

    // Warnings produces during execution.
    3: optional list<string> Warnings,
}
