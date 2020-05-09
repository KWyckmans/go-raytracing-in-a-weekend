package main

type HitableList struct {
	list []Hitable
}

func (h HitableList) hit(r Ray, tMin float32, tMax float32, rec *HitRecord) bool {
	var tempRec HitRecord
	var hitAny bool = false
	var closestSoFar float32 = tMax

	for _, v := range h.list {
		if v.hit(r, tMin, closestSoFar, &tempRec) {
			hitAny = true
			closestSoFar = tempRec.t
			*rec = tempRec
		}
	}

	return hitAny
}
