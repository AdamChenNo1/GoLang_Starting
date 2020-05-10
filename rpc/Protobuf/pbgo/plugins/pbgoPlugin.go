package pbgo

import (
	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

type pbgoPlugin struct {
	*generator.Generator
}

func (p *pbgoPlugin) Name() string {
	return "pbgo"
}

func (p *pbgoPlugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *pbgoPlugin) Generate(file *generator.FileDescriptor) {
	for _, svc := range file.Service {
		for _, m := range svc.Method {
			httpRule := p.getServiceMethod(m)
		}
	}
}

func (p *pbgoPlugin) getServiceMethod(m *descriptor.MethodDescriptorProto) *pbgo.HttpRule {
	if m.Options != nil && proto.HasExtension(m.options, pbgo.E_RestApi) {
		ext, _ := proto.GetExtension(m.Options, pbgo.E_RestApi)
		if ext != nil {
			if x, _ := ext.(*pbgo.HttpRule); x != nil {
				return x
			}
		}
	}
	return nil
}

func init() {
	generator.RegisterPlugin(new(pbgoPlugin))
}
