package simstring

type SizeWordMap struct {
	sizeMap map[int]map[string]struct{}
}

func MakeSizeWordMap() *SizeWordMap {
	return &SizeWordMap{
		sizeMap: make(map[int]map[string]struct{}),
	}
}

func (m *SizeWordMap) Add(size int, w string) {
	mapping, ok := m.sizeMap[size]
	if !ok {
		mapping = make(map[string]struct{})
		m.sizeMap[size] = mapping
	}
	mapping[w] = struct{}{}
}

func (m *SizeWordMap) Lookup(size int) (map[string]struct{}, bool) {
	mapping, ok := m.sizeMap[size]
	return mapping, ok

}

type SizeFeatureWordMap struct {
	sizeMap map[int]map[string]map[string]struct{}
}

func MakeSizeFeatureWordMap() *SizeFeatureWordMap {
	return &SizeFeatureWordMap{
		sizeMap: make(map[int]map[string]map[string]struct{}),
	}
}

func (m *SizeFeatureWordMap) Add(size int, f string, w string) {
	fmapping, ok := m.sizeMap[size]
	if !ok {
		fmapping = make(map[string]map[string]struct{})
		m.sizeMap[size] = fmapping
	}
	wmapping, ok := fmapping[f]
	if !ok {
		wmapping = make(map[string]struct{})
		fmapping[f] = wmapping
	}
	wmapping[w] = struct{}{}
}

func (m *SizeFeatureWordMap) Lookup(size int, f string) (map[string]struct{}, bool) {
	fmapping, ok := m.sizeMap[size]
	if !ok {
		return nil, ok
	}
	wmapping, ok := fmapping[f]
	return wmapping, ok
}

type DB interface {
	Add(string)
	Lookup(int, string) (map[string]struct{}, bool)
	Extract(string) []string
}

type OnMemoryDB struct {
	words          map[string]struct{}
	sizeFeatureMap *SizeFeatureWordMap
	maxFeatureSize int
	minFeatureSize int
	FeatureExtractor
}

func MakeOnMemoryDB(e FeatureExtractor) *OnMemoryDB {
	return &OnMemoryDB{
		words:            make(map[string]struct{}),
		sizeFeatureMap:   MakeSizeFeatureWordMap(),
		FeatureExtractor: e,
	}
}

func (d *OnMemoryDB) Add(w string) {
	if _, ok := d.words[w]; ok {
		return
	}
	features := d.Extract(w)
	size := len(features)
	d.updateMaxMin(size)
	d.words[w] = struct{}{}
	for _, f := range features {
		d.sizeFeatureMap.Add(size, f, w)
	}
}

func (d *OnMemoryDB) Lookup(size int, f string) (map[string]struct{}, bool) {
	if d.minFeatureSize <= size && size <= d.maxFeatureSize {
		mapping, ok := d.sizeFeatureMap.Lookup(size, f)
		return mapping, ok
	} else {
		return nil, false
	}

}

func (d *OnMemoryDB) updateMaxMin(size int) {
	if d.maxFeatureSize < size {
		d.maxFeatureSize = size
	}
	if d.minFeatureSize > size {
		d.minFeatureSize = size
	}
}
