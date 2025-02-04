/*
Generates OpenAPI v3.x.x from proto files.
*/
package main

import (
	"flag"
	"fmt"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/go-faster/errors"

	"github.com/ogen-go/protoc-gen-oas/internal/gen"
)

func run() error {
	set := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	openapi := set.String("openapi", "3.1.0", "OpenAPI version")
	title := set.String("title", "", "Title")
	description := set.String("description", "", "Description")
	version := set.String("version", "", "Version")
	indent := set.Int("indent", 2, "Indent")
	filename := set.String("filename", "openapi", "Filename")

	if err := set.Parse(os.Args[1:]); err != nil {
		return errors.Wrap(err, "parse args")
	}

	opts := protogen.Options{
		ParamFunc: set.Set,
	}

	p := func(plugin *protogen.Plugin) error {
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		g, err := gen.NewGenerator(
			plugin.Files,
			gen.WithSpecOpenAPI(*openapi),
			gen.WithSpecInfoTitle(*title),
			gen.WithSpecInfoDescription(*description),
			gen.WithSpecInfoVersion(*version),
			gen.WithIndent(*indent),
		)
		if err != nil {
			return err
		}

		openAPI, err := g.YAML()
		if err != nil {
			return err
		}

		bytes := make([]byte, 0)
		bytes = append(bytes, []byte("# generated by protoc-gen-oas. DO NOT EDIT\r\n\r\n")...)
		bytes = append(bytes, openAPI...)

		gf := plugin.NewGeneratedFile(fmt.Sprintf("%s.yaml", *filename), "")
		if _, err := gf.Write(bytes); err != nil {
			return err
		}

		return nil
	}

	opts.Run(p)

	return nil
}

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
