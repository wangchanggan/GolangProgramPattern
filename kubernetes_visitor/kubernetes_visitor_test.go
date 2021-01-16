package kubernetes_visitor

import "testing"

func TestKubernetesVisitor(t *testing.T) {
	info := Info{}
	var v Visitor = &info
	v = LogVisitor{v}
	v = NameVisitor{v}
	v = OtherThingsVisitor{v}

	loadFile := func(info *Info, err error) error {
		info.Name = "Hao Chen"
		info.Namespace = "MegaEase"
		info.OtherThings = "We are running as remote team."
		return nil
	}
	v.Visit(loadFile)

	/*LogVisitor() before call function
	NameVisitor() before call function
	OtherThingsVisitor() before call function
	==> OtherThings=We are running as remote team.
		OtherThingsVisitor() after call function
	==> Name=Hao Chen, NameSpace=MegaEase
	NameVisitor() after call function
	LogVisitor() after call function*/
}
