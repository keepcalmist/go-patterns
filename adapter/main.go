package main

func main() {
	gun := &Gun{}
	normalBullet := &NormalBullet{}

	gun.Shoot(normalBullet) //ok

	annormalBullet := &BigBullet{}
	bulletAdapter := &BigBulletAdapter{annormalBullet}

	gun.Shoot(bulletAdapter) //ok
}
