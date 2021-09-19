package main

type Gun struct {
}

type Bullet interface {
	InsertInto14mmGun()
}

func (g Gun) Shoot(bullet Bullet) {
	bullet.InsertInto14mmGun()
	//do something after insert ammo
}
