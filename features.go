package feature

import (
	"github.com/FrameworkOSS/portal/features/wires/wire"
)

func NewFeatureBinding(f Feature) (fb *FeatureBinding) {
	fb = new(FeatureBinding)
	fb.binding = f
	if f != nil {
		fb.
			SetID(f.ID()).
			SetName(f.Name()).
			SetAuthors(f.Authors()...).
			SetDescription(f.Description()).
			SetVersion(f.Version())
	}
	return
}

func NewFeatureBindingBytes(p []byte) (fb *FeatureBinding) {
	w := wire.NewWire(p)
	defer w.Close()

	id := w.GetValues(0)
	name := w.GetValues(1)
	authors := w.GetValues(2)
	description := w.GetValues(3)
	version := w.GetValues(4)

	fb = NewFeatureBinding(nil).
		SetID(string(id[0])).
		SetName(string(name[0])).
		SetDescription(string(description[0])).
		SetVersion(string(version[0]))

	if l := len(authors); l > 0 {
		a := make([]string, l)
		for i := 0; i < l; i++ {
			a[i] = string(authors[i])
		}
		fb.SetAuthors(a...)
	}

	return
}

func (fb *FeatureBinding) Bytes() []byte {
	w := wire.NewWire()
	defer w.Close()

	w.AddField(0, []byte(fb.id))
	w.AddField(1, []byte(fb.name))
	for i := 0; i < len(fb.authors); i++ {
		w.AddValues(2, []byte(fb.authors[i]))
	}
	w.AddField(3, []byte(fb.description))
	w.AddField(4, []byte(fb.version))
	return w.Bytes()
}

func (fb *FeatureBinding) API() int {
	if fb.binding != nil {
		return fb.binding.API()
	}
	return 0
}
func (fb *FeatureBinding) ID() string {
	return fb.id
}
func (fb *FeatureBinding) Name() string {
	return fb.name
}
func (fb *FeatureBinding) Authors() []string {
	return fb.authors
}
func (fb *FeatureBinding) Description() string {
	return fb.description
}
func (fb *FeatureBinding) Version() string {
	return fb.version
}
func (fb *FeatureBinding) Open() error {
	return nil
}
func (fb *FeatureBinding) Close() ([]error, bool) {
	return nil, false
}
func (fb *FeatureBinding) Input(e *Event) error {
	return fb.binding.Input(e)
}
func (fb *FeatureBinding) Output() (*Event, error) {
	if fb.binding == nil {
		return nil, nil
	}
	return fb.binding.Output()
}

func (fb *FeatureBinding) SetID(id string) *FeatureBinding {
	fb.id = id
	return fb
}

func (fb *FeatureBinding) SetName(name string) *FeatureBinding {
	fb.name = name
	return fb
}

func (fb *FeatureBinding) SetAuthors(authors ...string) *FeatureBinding {
	fb.authors = authors
	return fb
}

func (fb *FeatureBinding) SetDescription(description string) *FeatureBinding {
	fb.description = description
	return fb
}

func (fb *FeatureBinding) SetVersion(version string) *FeatureBinding {
	fb.version = version
	return fb
}

func (fb *FeatureBinding) SetBinding(f Feature) *FeatureBinding {
	fb.binding = f
	return fb
}

func (fb *FeatureBinding) GetBinding() Feature {
	return fb.binding
}

func (fb *FeatureBinding) CloneBinding() *FeatureBinding {
	return fb.
		SetID(fb.binding.ID()).
		SetName(fb.binding.Name()).
		SetAuthors(fb.binding.Authors()...).
		SetDescription(fb.binding.Description()).
		SetVersion(fb.binding.Version())
}
