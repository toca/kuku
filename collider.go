package main

import (
	"kuku/models"
)

type Collider struct {
	activeObjects     []models.Object
	motionlessObjects []models.Object
}

func NewCollider() *Collider {
	a := make([]models.Object, 0)
	m := make([]models.Object, 0)
	return &Collider{a, m}
}
func (this *Collider) AddActive(o models.Object) {
	this.activeObjects = append(this.activeObjects, o)
}

func (this *Collider) AddMotionless(o models.Object) {
	this.motionlessObjects = append(this.motionlessObjects, o)
}

func (this *Collider) RemoveActive(o *models.Object) {
	// TODO impl
}

func (this *Collider) RemoveMotionless(o *models.Object) {
	// TODO impl
}

func (this *Collider) Detect() {
	for ai, _ := range this.activeObjects {
		for mi, _ := range this.motionlessObjects {
			if this.motionlessObjects[mi].HitTest(this.activeObjects[ai]) {
				this.activeObjects[ai].Affect(this.motionlessObjects[mi])
				this.motionlessObjects[mi].Affect(this.activeObjects[ai])
			}
		}
	}
}
