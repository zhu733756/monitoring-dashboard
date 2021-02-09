// +build !ignore_autogenerated

/*
Copyright 2020 The KubeSphere authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package panels

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarGaugeOptions) DeepCopyInto(out *BarGaugeOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarGaugeOptions.
func (in *BarGaugeOptions) DeepCopy() *BarGaugeOptions {
	if in == nil {
		return nil
	}
	out := new(BarGaugeOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BarGaugePanel) DeepCopyInto(out *BarGaugePanel) {
	*out = *in
	out.Options = in.Options
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BarGaugePanel.
func (in *BarGaugePanel) DeepCopy() *BarGaugePanel {
	if in == nil {
		return nil
	}
	out := new(BarGaugePanel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gauge) DeepCopyInto(out *Gauge) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gauge.
func (in *Gauge) DeepCopy() *Gauge {
	if in == nil {
		return nil
	}
	out := new(Gauge)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Graph) DeepCopyInto(out *Graph) {
	*out = *in
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]Target, len(*in))
		copy(*out, *in)
	}
	if in.Colors != nil {
		in, out := &in.Colors, &out.Colors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Yaxes != nil {
		in, out := &in.Yaxes, &out.Yaxes
		*out = make([]Yaxis, len(*in))
		copy(*out, *in)
	}
	if in.Legend != nil {
		in, out := &in.Legend, &out.Legend
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Graph.
func (in *Graph) DeepCopy() *Graph {
	if in == nil {
		return nil
	}
	out := new(Graph)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Panel) DeepCopyInto(out *Panel) {
	*out = *in
	out.Options = in.Options
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]Target, len(*in))
		copy(*out, *in)
	}
	if in.Colors != nil {
		in, out := &in.Colors, &out.Colors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Color != nil {
		in, out := &in.Color, &out.Color
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Legend != nil {
		in, out := &in.Legend, &out.Legend
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HiddenColumns != nil {
		in, out := &in.HiddenColumns, &out.HiddenColumns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TimeSeriesAggregations != nil {
		in, out := &in.TimeSeriesAggregations, &out.TimeSeriesAggregations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Yaxes != nil {
		in, out := &in.Yaxes, &out.Yaxes
		*out = make([]Yaxis, len(*in))
		copy(*out, *in)
	}
	out.Gauge = in.Gauge
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Panel.
func (in *Panel) DeepCopy() *Panel {
	if in == nil {
		return nil
	}
	out := new(Panel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SingleStat) DeepCopyInto(out *SingleStat) {
	*out = *in
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]Target, len(*in))
		copy(*out, *in)
	}
	if in.Colors != nil {
		in, out := &in.Colors, &out.Colors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Color != nil {
		in, out := &in.Color, &out.Color
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Gauge = in.Gauge
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SingleStat.
func (in *SingleStat) DeepCopy() *SingleStat {
	if in == nil {
		return nil
	}
	out := new(SingleStat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Table) DeepCopyInto(out *Table) {
	*out = *in
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]Target, len(*in))
		copy(*out, *in)
	}
	if in.HiddenColumns != nil {
		in, out := &in.HiddenColumns, &out.HiddenColumns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TimeSeriesAggregations != nil {
		in, out := &in.TimeSeriesAggregations, &out.TimeSeriesAggregations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Table.
func (in *Table) DeepCopy() *Table {
	if in == nil {
		return nil
	}
	out := new(Table)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Target) DeepCopyInto(out *Target) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Target.
func (in *Target) DeepCopy() *Target {
	if in == nil {
		return nil
	}
	out := new(Target)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Text) DeepCopyInto(out *Text) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Text.
func (in *Text) DeepCopy() *Text {
	if in == nil {
		return nil
	}
	out := new(Text)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Yaxis) DeepCopyInto(out *Yaxis) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Yaxis.
func (in *Yaxis) DeepCopy() *Yaxis {
	if in == nil {
		return nil
	}
	out := new(Yaxis)
	in.DeepCopyInto(out)
	return out
}
