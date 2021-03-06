# Dear ImGui for Go

[![Go Report Card](https://goreportcard.com/badge/github.com/inkyblackness/imgui-go)](https://goreportcard.com/report/github.com/inkyblackness/imgui-go)

This library is a [Go](https://www.golang.org) wrapper for [Dear ImGui](https://github.com/ocornut/imgui).

At the moment, this wrapper is a special-purpose wrapper for use within InkyBlackness.
However, it is self-contained and can be used for other purposes as well.

This wrapper is
* hand-crafted, for Go
* documented
* versioned
* with a ported example using [GLFW3](https://github.com/go-gl/glfw) and [OpenGL3](https://github.com/go-gl/gl).

![Screenshot from example](_examples/opengl3_example/screenshot.png)

## API naming

Names of types and functions follow closely those of **Dear ImGui**.

For functions that have optional parameters, the following schema is applied:
* There is the "verbose" variant, followed by the letter `V`, such as `ButtonV(id string, size Vec2) bool`
* Next to it there is the "idiomatic" variant, without any optional parameter, such as `Button(id string) bool`.
* The idiomatic variant calls the verbose variant with the default values for the optional parameters.
Functions that don't have optional parameters don't come in a verbose variant.

The **Dear ImGui** functions `IO()` and `Style()` have been renamed to be `CurrentIO()` and `CurrentStyle()`.
This was done because their returned types have the same name, causing a name clash.
With the `Current` prefix, they also better describe what they return.  

## API philosophy
This library does not intend to export all the functions of the wrapped ImGui. The following filter applies as a rule of thumb:
* Functions marked as "obsolete" are not available. (The corresponding C code isn't even compiled - disabled by define)
* "Shortcut" Functions, which combine language features and/or other ImGui functions, are not available. Prime example are the Text*() functions for instance: Text formatting should be done with fmt.Sprintf(), and style formatting with the corresponding Push/Pop functions.
* Functions that are not needed by InkyBlackness are ignored. This doesn't mean that they can't be in the wrapper, they are simply not a priority. Feel free to propose an implementation or make a pull request, respecting the previous points :)

## Version philosophy
This library does not mirror the versions of the wrapped ImGui. The semantic versioning of this wrapper is defined as:
* Major changes: (Breaking) changes in API or behaviour. Typically done through changes in ImGui.
* Minor changes: Extensions in API. Typically done through small version increments of ImGui and/or exposing further features in a compatible way.
* Patch changes: Bug fixes - either in the wrapper or the wrapped ImGui, given that the API & behaviour remains the same.

At the moment, this library uses version [1.61](https://github.com/ocornut/imgui/releases/tag/v1.61) of ImGui.

## Alternatives

Before this project was created, the following alternatives were considered - and ignored:
* [kdrag0n/go-imgui](https://github.com/kdrag0n/go-imgui). Reasons for dismissal at time of decision:
  * Auto-generated bloat, which doesn't help
  * Was using old API (1.5x)
  * Did not compile (Issues [1](https://github.com/kdrag0n/go-imgui/issues/1) and [3](https://github.com/kdrag0n/go-imgui/issues/3))
  * Project appeared to be abandoned
* [Extrawurst/cimgui](https://github.com/Extrawurst/cimgui). Reasons for dismissal at time of decision:
  * Was using old API (1.5x), 1.6x was attempted
  * Apparently semi-exposed the C++ API, especially through the structures
  * Adding this adds another dependency


## License

The project is available under the terms of the **New BSD License** (see LICENSE file).
The licenses of included sources are stored in the **_licenses** folder.
