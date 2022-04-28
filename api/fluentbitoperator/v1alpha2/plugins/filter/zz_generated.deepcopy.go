//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*

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

package filter

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	if in.KeyDoesNotExist != nil {
		in, out := &in.KeyDoesNotExist, &out.KeyDoesNotExist
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.KeyValueEquals != nil {
		in, out := &in.KeyValueEquals, &out.KeyValueEquals
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.KeyValueDoesNotEqual != nil {
		in, out := &in.KeyValueDoesNotEqual, &out.KeyValueDoesNotEqual
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.KeyValueMatches != nil {
		in, out := &in.KeyValueMatches, &out.KeyValueMatches
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.KeyValueDoesNotMatch != nil {
		in, out := &in.KeyValueDoesNotMatch, &out.KeyValueDoesNotMatch
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MatchingKeysHaveMatchingValues != nil {
		in, out := &in.MatchingKeysHaveMatchingValues, &out.MatchingKeysHaveMatchingValues
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MatchingKeysDoNotHaveMatchingValues != nil {
		in, out := &in.MatchingKeysDoNotHaveMatchingValues, &out.MatchingKeysDoNotHaveMatchingValues
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Grep) DeepCopyInto(out *Grep) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Grep.
func (in *Grep) DeepCopy() *Grep {
	if in == nil {
		return nil
	}
	out := new(Grep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kubernetes) DeepCopyInto(out *Kubernetes) {
	*out = *in
	if in.MergeLog != nil {
		in, out := &in.MergeLog, &out.MergeLog
		*out = new(bool)
		**out = **in
	}
	if in.MergeLogTrim != nil {
		in, out := &in.MergeLogTrim, &out.MergeLogTrim
		*out = new(bool)
		**out = **in
	}
	if in.KeepLog != nil {
		in, out := &in.KeepLog, &out.KeepLog
		*out = new(bool)
		**out = **in
	}
	if in.TLSDebug != nil {
		in, out := &in.TLSDebug, &out.TLSDebug
		*out = new(int32)
		**out = **in
	}
	if in.TLSVerify != nil {
		in, out := &in.TLSVerify, &out.TLSVerify
		*out = new(bool)
		**out = **in
	}
	if in.UseJournal != nil {
		in, out := &in.UseJournal, &out.UseJournal
		*out = new(bool)
		**out = **in
	}
	if in.K8SLoggingParser != nil {
		in, out := &in.K8SLoggingParser, &out.K8SLoggingParser
		*out = new(bool)
		**out = **in
	}
	if in.K8SLoggingExclude != nil {
		in, out := &in.K8SLoggingExclude, &out.K8SLoggingExclude
		*out = new(bool)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = new(bool)
		**out = **in
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = new(bool)
		**out = **in
	}
	if in.DummyMeta != nil {
		in, out := &in.DummyMeta, &out.DummyMeta
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kubernetes.
func (in *Kubernetes) DeepCopy() *Kubernetes {
	if in == nil {
		return nil
	}
	out := new(Kubernetes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Lua) DeepCopyInto(out *Lua) {
	*out = *in
	in.Script.DeepCopyInto(&out.Script)
	if in.TypeIntKey != nil {
		in, out := &in.TypeIntKey, &out.TypeIntKey
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ProtectedMode != nil {
		in, out := &in.ProtectedMode, &out.ProtectedMode
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Lua.
func (in *Lua) DeepCopy() *Lua {
	if in == nil {
		return nil
	}
	out := new(Lua)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Modify) DeepCopyInto(out *Modify) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Modify.
func (in *Modify) DeepCopy() *Modify {
	if in == nil {
		return nil
	}
	out := new(Modify)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Nest) DeepCopyInto(out *Nest) {
	*out = *in
	if in.Wildcard != nil {
		in, out := &in.Wildcard, &out.Wildcard
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Nest.
func (in *Nest) DeepCopy() *Nest {
	if in == nil {
		return nil
	}
	out := new(Nest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parser) DeepCopyInto(out *Parser) {
	*out = *in
	if in.PreserveKey != nil {
		in, out := &in.PreserveKey, &out.PreserveKey
		*out = new(bool)
		**out = **in
	}
	if in.ReserveData != nil {
		in, out := &in.ReserveData, &out.ReserveData
		*out = new(bool)
		**out = **in
	}
	if in.UnescapeKey != nil {
		in, out := &in.UnescapeKey, &out.UnescapeKey
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parser.
func (in *Parser) DeepCopy() *Parser {
	if in == nil {
		return nil
	}
	out := new(Parser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordModifier) DeepCopyInto(out *RecordModifier) {
	*out = *in
	if in.Records != nil {
		in, out := &in.Records, &out.Records
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RemoveKeys != nil {
		in, out := &in.RemoveKeys, &out.RemoveKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WhitelistKeys != nil {
		in, out := &in.WhitelistKeys, &out.WhitelistKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordModifier.
func (in *RecordModifier) DeepCopy() *RecordModifier {
	if in == nil {
		return nil
	}
	out := new(RecordModifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RewriteTag) DeepCopyInto(out *RewriteTag) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RewriteTag.
func (in *RewriteTag) DeepCopy() *RewriteTag {
	if in == nil {
		return nil
	}
	out := new(RewriteTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	if in.Set != nil {
		in, out := &in.Set, &out.Set
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Add != nil {
		in, out := &in.Add, &out.Add
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Rename != nil {
		in, out := &in.Rename, &out.Rename
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HardRename != nil {
		in, out := &in.HardRename, &out.HardRename
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Copy != nil {
		in, out := &in.Copy, &out.Copy
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HardCopy != nil {
		in, out := &in.HardCopy, &out.HardCopy
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Throttle) DeepCopyInto(out *Throttle) {
	*out = *in
	if in.Rate != nil {
		in, out := &in.Rate, &out.Rate
		*out = new(int64)
		**out = **in
	}
	if in.Window != nil {
		in, out := &in.Window, &out.Window
		*out = new(int64)
		**out = **in
	}
	if in.PrintStatus != nil {
		in, out := &in.PrintStatus, &out.PrintStatus
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Throttle.
func (in *Throttle) DeepCopy() *Throttle {
	if in == nil {
		return nil
	}
	out := new(Throttle)
	in.DeepCopyInto(out)
	return out
}
