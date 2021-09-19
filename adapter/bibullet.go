package main

type BigBullet struct {
}

func (b *BigBullet) InsertInto25mmGun() {

}

type BigBulletAdapter struct {
	*BigBullet
}

func (b *BigBullet) InsertInto14mmGun() {

}
