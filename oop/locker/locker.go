package locker

import (
	"errors"
	"sort"
)

// A uuid is a 128 bit unique identifier of information
type uuid [16]byte

// A UUIDGenerator can be used to generate uuids
type UUIDGenerator interface {
	GetUUID() uuid
}

// A _package represents a package that will be stored in the locker
type _package struct {
	id uuid
	size int
	Content interface{}
}

// A space may hold a _package of equal or lesser size
type space struct {
	_package *_package
	size int
}

// A locker contains spaces of varying sizes, each of which may or may not contain _packages
// It also maintains a mapping of uuids to spaces so that it can open a space when a valid code is provided
type locker struct {
	spaces        []space
	idsToSpaces map[uuid]*space
}

// A spaceTuple holds a space size and a number of spaces, which is useful when constructing a locker
type spaceTuple struct {
	size, numSpaces int
}

// TODO replace with an actual UUIDGenerator implementation
func getUUIDGenerator() UUIDGenerator {
	return nil
}

// Creates a new _package of the given size containing the given content
func NewPackage(size int, content interface{}) *_package {
	return &_package{id: getUUIDGenerator().GetUUID(), size: size, Content: content}
}

// Creates a new locker given a map whose keys represent space sizes and whose values represent the number of desired
// spaces of that size e.g. a mapping of 2:10 would mean "this locker should have 10 spaces of size 2"
func NewLocker(desiredSpaces map[int]int) *locker {
	spaces := make([]space, len(desiredSpaces))
	tuples := make([]spaceTuple, len(desiredSpaces))
	for size, numSpaces := range desiredSpaces {
		tuples = append(tuples, spaceTuple{size, numSpaces})
	}
	sort.Slice(tuples, func (i int, j int) bool { return tuples[i].size < tuples[j].size })
	for _, tuple := range tuples {
		for i := 0; i < tuple.numSpaces; i++ {
			spaces = append(spaces, space{nil, tuple.size})
		}
	}
	return &locker{spaces, make(map[uuid]*space)}
}

// getPackage retrieves a _package from a space then makes the space empty
func (s *space) getPackage() *_package {
	ret := s._package
	s._package = nil
	return ret
}

// addPackage places a _package into a space if the space is empty and of sufficient size
func (s *space) addPackage(p *_package) error {
	if s._package != nil {
		return errors.New("space already contains a package")
	} else if s.size < p.size {
		return errors.New("space is not large enough to hold the package")
	}
	s._package = p
	return nil
}

// closestFit returns the space that is closest in size to the given size
// a and b must either be nil or spaces with size >= the given size
func closestFit(a, b *space, size int) *space {
	if a == nil && b == nil {
		return nil
	} else if b == nil {
		return a
	} else if a == nil {
		return b
	}
	if a.size - size < b.size - size {
		return a
	}
	return b
}

// bSearch finds an empty space that is closest in size to the given size, or returns nil if none is found
func bSearch(spaces []space, lo int, hi int, size int) *space {
	if lo > hi {
		return nil
	}
	mid := (lo + hi) / 2
	var curBestSpace *space
	var nextBestSpace *space
	if spaces[mid].size == size {
		if spaces[mid]._package == nil {
			return &spaces[mid]
		} else {
			leftBestSpace := bSearch(spaces, lo, mid-1, size)
			rightBestSpace := bSearch(spaces, mid+1, hi, size)
			return closestFit(leftBestSpace, rightBestSpace, size)
		}
	} else if spaces[mid].size > size {
		if spaces[mid]._package == nil {
			curBestSpace = &spaces[mid]
		}
		nextBestSpace = bSearch(spaces, lo, mid-1, size)
		return closestFit(curBestSpace, nextBestSpace, size)
	} else {
		return bSearch(spaces, mid+1, hi, size)
	}
}

// AddPackage puts a Package into the locker, provided there is an open space of sufficient size
func (l *locker) AddPackage(p *_package) error {
	space := bSearch(l.spaces, 0, len(l.spaces), p.size)
	if space == nil {
		return errors.New("no available space for package")
	}
	err := space.addPackage(p)
	if err != nil {
		panic("could not add package to seemingly valid space")
	}
	l.idsToSpaces[p.id] = space
	return nil
}

// RetrievePackage retrieves a Package from the locker, given that the provided code is valid
func (l *locker) RetrievePackage(id uuid) (*_package, error) {
	space, ok := l.idsToSpaces[id]
	if !ok {
		return nil, errors.New("invalid code")
	}
	p := space.getPackage()
	delete(l.idsToSpaces, id)
	return p, nil
}