package goasteroids

import (
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/solarlune/resolv"

	"go-asteroids/assets"
)

const (
	rotationPerSecond    = math.Pi
	maxAcceleration      = 8.0
	ScreenWidth          = 1280 // The width of the screen. We use a 16/9 aspect ratio.
	ScreenHeight         = 720  // The height of the screen.
	shootCoolDown        = time.Millisecond * 150
	burstCoolDown        = time.Millisecond * 500
	laserSpawnOffSet     = 50.0
	maxShotsPerBurts     = 3
	dyingAnimationAmount = 50 * time.Millisecond
	numberOfLives        = 3
	numberOfShields      = 3
	shieldDuration       = time.Second * 6
	hyperSpaceCooldown   = time.Second * 10
)

var (
	curAcceleration float64 // We use this to gradually increase acceleration.
	shotsFired      = 0     // A counter to keep track of max shots per burst.
)

type Player struct {
	game                *GameScene    // The current game scene.
	sprite              *ebiten.Image // The player's sprite.
	rotation            float64       // The current player's rotation.
	position            Vector        // Where is the player on the screen.
	playerVelocity      float64
	playerObj           *resolv.Circle
	shootCoolDown       *Timer
	burstCoolDown       *Timer
	isShielded          bool
	isDying             bool
	isDead              bool
	dyingTimer          *Timer
	dyingCounter        int
	livesRemaining      int
	lifeIndicators      []*LifeIndicator
	shieldTimer         *Timer
	shieldsRemaining    int
	shieldIndicators    []*ShieldIndicator
	hyperspaceIndicator *HyperspaceIndicator
	hyperSpaceTimer     *Timer
}

func NewPlayer(game *GameScene) *Player {
	sprite := assets.PlayerSprite

	// Center player on screen.
	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos := Vector{
		X: ScreenWidth/2 - halfW,
		Y: ScreenHeight/2 - halfH,
	}

	// Create a resolv object.
	playerObj := resolv.NewCircle(pos.X, pos.Y, float64(sprite.Bounds().Dx()/2))

	var lifeIndicators []*LifeIndicator
	var xPosition = 20.0

	for i := 0; i < numberOfLives; i++ {
		li := NewLifeIndicator(Vector{X: xPosition, Y: 20})
		lifeIndicators = append(lifeIndicators, li)
		xPosition += 50.0
	}

	var shieldIndicators []*ShieldIndicator
	xPosition = 45.0

	for i := 0; i < numberOfShields; i++ {
		si := NewShieldIndicator(Vector{X: xPosition, Y: 60})
		shieldIndicators = append(shieldIndicators, si)
		xPosition += 50.0
	}

	p := &Player{
		sprite:              sprite,
		game:                game,
		position:            pos,
		playerObj:           playerObj,
		shootCoolDown:       NewTimer(shootCoolDown),
		burstCoolDown:       NewTimer(burstCoolDown),
		isShielded:          false,
		isDying:             false,
		isDead:              false,
		dyingTimer:          NewTimer(dyingAnimationAmount),
		dyingCounter:        0,
		livesRemaining:      numberOfLives,
		lifeIndicators:      lifeIndicators,
		shieldsRemaining:    numberOfShields,
		shieldIndicators:    shieldIndicators,
		hyperspaceIndicator: NewHyperspaceIndicator(Vector{X: 37.0, Y: 95.0}),
		hyperSpaceTimer:     nil,
	}

	p.playerObj.SetPosition(pos.X, pos.Y)
	p.playerObj.Tags().Set(TagPlayer)

	return p
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, op)
}

func (p *Player) Update() {
	speed := rotationPerSecond / float64(ebiten.TPS())

	p.isPlayerDead()

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += speed
	}

	p.accelerate()

	p.useShield()

	p.isDoneAccelerating()

	p.reverse()
	p.isDoneReversing()

	p.updateExhaustSprite()

	p.playerObj.SetPosition(p.position.X, p.position.Y)

	p.burstCoolDown.Update()
	p.shootCoolDown.Update()

	p.fireLasers()

	p.hyperSpace()

	if p.hyperSpaceTimer != nil {
		p.hyperSpaceTimer.Update()
	}
}

func (p *Player) hyperSpace() {
	if ebiten.IsKeyPressed(ebiten.KeyH) && (p.hyperSpaceTimer == nil || p.hyperSpaceTimer.IsReady()) {
		var randX, randY int
		for {
			randX = rand.Intn(ScreenWidth)
			randY = rand.Intn(ScreenHeight)

			collision := p.game.checkCollision(p.playerObj, nil)
			if !collision {
				break
			}

			p.position.X = float64(randX)
			p.position.Y = float64(randY)

			if p.hyperSpaceTimer == nil {
				p.hyperSpaceTimer = NewTimer(hyperSpaceCooldown)
			}

			p.hyperSpaceTimer.Reset()
		}
	}
}

func (p *Player) useShield() {
	if ebiten.IsKeyPressed(ebiten.KeyS) && !p.isShielded && p.shieldsRemaining > 0 {
		if !p.game.shieldsUpPlayer.IsPlaying() {
			_ = p.game.shieldsUpPlayer.Rewind()
			p.game.shieldsUpPlayer.Play()
		}

		p.isShielded = true
		p.shieldTimer = NewTimer(shieldDuration)
		p.game.shield = NewShield(Vector{}, p.rotation, p.game)
		p.shieldsRemaining--
		p.shieldIndicators = p.shieldIndicators[:len(p.shieldIndicators)-1]
	}

	if p.shieldTimer != nil && p.isShielded {
		p.shieldTimer.Update()
	}

	if p.shieldTimer != nil && p.shieldTimer.IsReady() {
		p.shieldTimer = nil
		p.isShielded = false
		p.game.space.Remove(p.game.shield.shieldObj)
		p.game.shield = nil
	}
}

func (p *Player) isPlayerDead() {
	if p.isDead {
		p.game.playerIsDead = true
	}
}

func (p *Player) fireLasers() {
	if p.burstCoolDown.IsReady() {
		if p.shootCoolDown.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace) {
			p.shootCoolDown.Reset()
			shotsFired++
			if shotsFired <= maxShotsPerBurts {
				bounds := p.sprite.Bounds()
				halfW := float64(bounds.Dx()) / 2
				halfH := float64(bounds.Dy()) / 2

				spawnPos := Vector{
					p.position.X + halfW + math.Sin(p.rotation)*laserSpawnOffSet,
					p.position.Y + halfH + math.Cos(p.rotation)*-laserSpawnOffSet,
				}

				p.game.laserCount++
				laser := NewLaser(spawnPos, p.rotation, p.game.laserCount, p.game)
				p.game.lasers[p.game.laserCount] = laser
				p.game.space.Add(laser.laserObj)

				switch shotsFired {
				case 1:
					if !p.game.laserOnePlayer.IsPlaying() {
						_ = p.game.laserOnePlayer.Rewind()
						p.game.laserOnePlayer.Play()
					}
				case 2:
					if !p.game.laserTwoPlayer.IsPlaying() {
						_ = p.game.laserTwoPlayer.Rewind()
						p.game.laserTwoPlayer.Play()
					}
				case 3:
					if !p.game.laserThreePlayer.IsPlaying() {
						_ = p.game.laserThreePlayer.Rewind()
						p.game.laserThreePlayer.Play()
					}
				}
			} else {
				p.burstCoolDown.Reset()
				shotsFired = 0
			}
		}
	}
}

func (p *Player) accelerate() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.keepOnScreen()

		if curAcceleration < maxAcceleration {
			curAcceleration = p.playerVelocity + 4
		}

		if curAcceleration >= 8 {
			curAcceleration = 8
		}

		p.playerVelocity = curAcceleration

		// Move in the direction we are pointing.
		dx := math.Sin(p.rotation) * curAcceleration
		dy := math.Cos(p.rotation) * -curAcceleration

		// Show exhaust.
		bounds := p.sprite.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfH := float64(bounds.Dy()) / 2

		// Where to spawn exhaust?
		spawnPos := Vector{
			p.position.X + halfW + math.Sin(p.rotation)*exhaustSpawnOffset,
			p.position.Y + halfH + math.Cos(p.rotation)*-exhaustSpawnOffset,
		}

		p.game.exhaust = NewExhaust(spawnPos, p.rotation+180.0*math.Pi/180.0)

		// Move the player on the screen.
		p.position.X += dx
		p.position.Y += dy

		if !p.game.thrustPlayer.IsPlaying() {
			_ = p.game.thrustPlayer.Rewind()
			p.game.thrustPlayer.Play()
		}
	}
}

func (p *Player) isDoneAccelerating() {
	if inpututil.IsKeyJustReleased(ebiten.KeyUp) {
		if p.game.thrustPlayer.IsPlaying() {
			p.game.thrustPlayer.Pause()
		}
	}
}

func (p *Player) reverse() {
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.keepOnScreen()

		dx := math.Sin(p.rotation) * -3
		dy := math.Cos(p.rotation) * 3

		bounds := p.sprite.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfH := float64(bounds.Dy()) / 2

		spawnPos := Vector{
			p.position.X + halfW + math.Sin(p.rotation)*-exhaustSpawnOffset,
			p.position.Y + halfH + math.Cos(p.rotation)*exhaustSpawnOffset,
		}

		p.game.exhaust = NewExhaust(spawnPos, p.rotation+180.0*math.Pi/180.0)

		p.position.X += dx
		p.position.Y += dy

		p.playerObj.SetPosition(p.position.X, p.position.Y)

		if !p.game.thrustPlayer.IsPlaying() {
			_ = p.game.thrustPlayer.Rewind()
			p.game.thrustPlayer.Play()
		}
	}
}

func (p *Player) isDoneReversing() {
	if inpututil.IsKeyJustReleased(ebiten.KeyDown) {
		if p.game.thrustPlayer.IsPlaying() {
			p.game.thrustPlayer.Pause()
		}
	}
}

func (p *Player) updateExhaustSprite() {
	if !ebiten.IsKeyPressed(ebiten.KeyUp) && !ebiten.IsKeyPressed(ebiten.KeyDown) && p.game.exhaust != nil {
		p.game.exhaust = nil
	}
}

func (p *Player) keepOnScreen() {
	if p.position.X >= float64(ScreenWidth) {
		p.position.X = 0
		p.playerObj.SetPosition(0, p.position.Y)
	}
	if p.position.X < 0 {
		p.position.X = ScreenWidth
		p.playerObj.SetPosition(ScreenWidth, p.position.Y)
	}
	if p.position.Y >= float64(ScreenHeight) {
		p.position.Y = 0
		p.playerObj.SetPosition(p.position.X, 0)
	}
	if p.position.Y < 0 {
		p.position.Y = ScreenHeight
		p.playerObj.SetPosition(p.position.X, ScreenHeight)
	}
}
