package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
}

func (b *Boid) calcAcceleration() Vector2D {
	upper, lower := b.position.AddV(viewRadius), b.position.AddV(-viewRadius)
	avgPosition, avgVelocity, separation := Vector2D{0, 0}, Vector2D{0, 0}, Vector2D{0, 0}
	count := 0.0
	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight); j++ {
			if otherBoidId := boidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.id {
				if dist := boids[otherBoidId].position.Distance(b.position); dist < viewRadius {
					count++
					avgVelocity = avgVelocity.Add(boids[otherBoidId].velocity)
					avgPosition = avgPosition.Add(boids[otherBoidId].position)
					separation = separation.Add(b.position.Subtract(boids[otherBoidId].position).DivisionV(dist))
				}
			}
		}
	}
	accel := Vector2D{b.borderBounce(b.position.x, screenWidth), b.borderBounce(b.position.y, screenHeight)}
	if count > 0 {
		avgPosition, avgVelocity = avgPosition.DivisionV(count), avgVelocity.DivisionV(count)
		accelAlignment := avgVelocity.Subtract(b.velocity).MultiplyV(adjRate)
		accelCohesion := avgPosition.Subtract(b.position).MultiplyV(adjRate)
		accelSeparation := separation.MultiplyV(adjRate)
		accel = accel.Add(accelAlignment).Add(accelCohesion).Add(accelSeparation)
	}
	return accel
}

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < viewRadius {
		return 1 / pos
	} else if pos > maxBorderPos-viewRadius {
		return 1 / (pos - maxBorderPos)
	}
	return 0
}

func (b *Boid) moveOne() {
	acceleration := b.calcAcceleration()
	b.velocity = b.velocity.Add(acceleration).Limit(-1, 1)
	boidMap[int(b.position.x)][int(b.position.y)] = -1
	b.position = b.position.Add(b.velocity)
	boidMap[int(b.position.x)][int(b.position.y)] = b.id
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(bid int) {
	b := Boid{
		position: Vector2D{rand.Float64() * screenWidth, rand.Float64() * screenHeight},
		velocity: Vector2D{(rand.Float64() * 2) - 1.0, (rand.Float64() * 2) - 1.0},
		id:       bid,
	}
	boids[bid] = &b
	boidMap[int(b.position.x)][int(b.position.y)] = b.id
	go b.start()
}

// package main
//
// import (
// 	"math"
// 	"math/rand"
// 	"time"
// )
//
// type Boid struct {
// 	position Vector2D
// 	velocity Vector2D
// 	id       int
// }
//
// func (b *Boid) calcAcceleration() Vector2D {
// 	upper, lower := b.position.AddV(viewRadius), b.position.AddV(-viewRadius)
// 	avgPosition, avgVelocity, separation := Vector2D{0, 0}, Vector2D{0, 0}, Vector2D{0, 0}
// 	count := 0.0
//
// 	// we are  aquiring(aquiere) the mutex, blocking other threads to move
// 	rwLock.RLock()
// 	// lower or 0 if lower < 0
// 	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth); i++ {
// 		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight); j++ {
// 			if otherBoidId := boidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.id {
// 				if dist := boids[otherBoidId].position.Distance(b.position); dist < viewRadius {
// 					count++
// 					avgVelocity = avgVelocity.Add(boids[otherBoidId].velocity)
// 					avgPosition = avgPosition.Add(boids[otherBoidId].position)
// 					separation = separation.Add(b.position.Subtract(boids[otherBoidId].position).DivisionV(dist))
// 				}
// 			}
// 		}
// 	}
// 	// we release the mutex
// 	rwLock.RUnlock()
//
// 	accel := Vector2D{
// 		b.borderBounce(b.position.x, screenWidth),
// 		b.borderBounce(b.position.y, screenHeight),
// 	}
// 	if count > 0 {
// 		avgPosition, avgVelocity = avgPosition.DivisionV(count), avgVelocity.DivisionV(count)
// 		accelAlignment := avgVelocity.Subtract(b.velocity).MultiplyV(adjRate)
// 		accelCohesion := avgPosition.Subtract(b.position).MultiplyV(adjRate)
// 		accelSeparation := separation.MultiplyV(adjRate)
// 		accel = accel.Add(accelAlignment).Add(accelCohesion).Add(accelSeparation)
// 	}
//
// 	return accel
// }
//
// func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
// 	if pos < viewRadius {
// 		return 1 / pos
// 	} else if pos > maxBorderPos-viewRadius {
// 		return 1 / (pos - maxBorderPos)
// 	}
//
// 	return 0
// }
//
// func (b *Boid) moveOne() {
// 	// remove b.calcAcceleration outside the lock
// 	// otherwise the thread will lock itself
// 	acceleration := b.calcAcceleration()
// 	rwLock.Lock()
// 	b.velocity = b.velocity.Add(acceleration).Limit(-1, 1)
// 	boidMap[int(b.position.x)][int(b.position.y)] = -1
// 	b.position = b.position.Add(b.velocity)
// 	boidMap[int(b.position.x)][int(b.position.y)] = b.id
// 	// next := b.position.Add(b.velocity)
// 	// if next.x >= screenWidth || next.x <= 0 {
// 	// 	b.velocity = Vector2D{-b.velocity.x, b.velocity.y}
// 	// }
// 	// if next.y >= screenHeight || next.y <= 0 {
// 	// 	b.velocity = Vector2D{b.velocity.x, -b.velocity.y}
// 	// }
// 	rwLock.Unlock()
// }
//
// func (b *Boid) start() {
// 	for {
// 		b.moveOne()
// 		time.Sleep(5 * time.Millisecond)
// 	}
// }
//
// func createBoid(bid int) {
// 	b := Boid{
// 		// rand.Float64 return a num between 0 & 1
// 		position: Vector2D{rand.Float64() * screenWidth, rand.Float64() * screenHeight},
// 		velocity: Vector2D{(rand.Float64() * 2) - 1.0, (rand.Float64() * 2) - 1.0},
// 		id:       bid,
// 	}
// 	boids[bid] = &b
// 	boidMap[int(b.position.x)][int(b.position.y)] = b.id
//
// 	go b.start()
// }