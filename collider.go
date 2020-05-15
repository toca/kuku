package main

import (
	"kuku/models"
)

type Collider struct {
	activeObjects     []models.Object
	motionlessObjects []models.Object
	// lastDetected      time.Time
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
	// const rate = 16 // 1000 / 60 (FPS)
	// d := time.Now().Sub(this.lastDetected)
	// if float64(d.Milliseconds()) < rate {
	// 	return
	// }
	for ai := 0; ai < len(this.activeObjects); ai++ {
		if this.activeObjects[ai].MarkedForDeath() {
			this.activeObjects = append(this.activeObjects[:ai], this.activeObjects[ai+1:]...)
			ai++
			if len(this.activeObjects) <= ai {
				break
			}
		}
		for mi := 0; mi < len(this.motionlessObjects); mi++ {
			if this.motionlessObjects[mi].MarkedForDeath() {
				this.motionlessObjects = append(this.motionlessObjects[:mi], this.motionlessObjects[mi+1:]...)
				mi++
				if len(this.motionlessObjects) <= mi {
					break
				}
			}
			if this.motionlessObjects[mi].HitTest(this.activeObjects[ai]) {
				this.activeObjects[ai].Affect(this.motionlessObjects[mi])
				this.motionlessObjects[mi].Affect(this.activeObjects[ai])
			}
		}
	}
	// this.lastDetected = this.lastDetected.Add(d)
}
