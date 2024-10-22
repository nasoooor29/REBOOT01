const ball = document.querySelector('#ball')
const GameContainer = document.querySelector('#game')
const padel = document.querySelector('#player')
const bricksContainer = document.querySelector('#bricksContainer')
const tutorialBox = document.querySelector('#tutorial')
const livesText = document.getElementById('livesText')
const gameOverBox = document.getElementById('gameover')
const countdown = document.getElementById('countdown')
const scoreText = document.getElementById('scoreText')
const pauseGame = document.getElementById('gamepause')
const pauseBox = document.getElementById('pauseInst')
const winBox = document.getElementById('win')
const scoreDisplay = document.querySelectorAll('.displayScore')

let ballRect = () => ball.getBoundingClientRect()
let gameRect = () => GameContainer.getBoundingClientRect()
let padelRect = () => padel.getBoundingClientRect()

let lives = 3
let score = 0
let brickTracker = 49
let gameOver = false
let win = false
let started = false
padelPosX = 0
let padelSpeed = 2.5
let speedX = 1.5
let speedY = -speedX
let countdownInterval
let targetTime = 180
let currentTime = targetTime + 1
let isPaused = false
let paused = false
let bricks = [...document.querySelectorAll('.brick')]

let mapp = {
  ArrowRight: false,
  ArrowLeft: false,
  s: false,
  r: false,
  p: false
}

createBricks()

function startTimer () {
  updateTimer()
  countdownInterval = setInterval(updateTimer, 1000)
}

function updateTimer () {
  currentTime--
  let minutes = 0
  for (let i = currentTime; i >= 60; i -= 60) {
    minutes++
  }
  let seconds = currentTime - minutes * 60
  minutes = minutes < 10 ? '0' + minutes : minutes
  seconds = seconds < 10 ? '0' + seconds : seconds

  document.getElementById('minutes').innerText = minutes
  document.getElementById('seconds').innerText = seconds

  if (minutes <= 0 && seconds <= 0) {
    timeOver()
  }
}

function pauseTimer () {
  if (!isPaused) {
    clearInterval(countdownInterval)
    isPaused = true
  }
}

function resumeTimer () {
  if (isPaused) {
    countdownInterval = setInterval(updateTimer, 1000)
    isPaused = false
  }
}

document.addEventListener('keydown', event => {
  if (event.key.toLowerCase() === 'p' && !gameOver && started) {
    if (paused) {
      paused = false
      pauseGame.style.display = 'none'
    } else {
      paused = true
      pauseGame.style.display = 'block'
    }
    if (isPaused) {
      resumeTimer()
    } else {
      pauseTimer()
    }
  }
})

document.addEventListener('keydown', event => {
  if (mapp.hasOwnProperty(event.key) && event.key !== 's') {
    mapp[event.key] = true
  }
  if (event.key === 's' && !gameOver) {
    mapp.s = true
    started = false
    tutorialBox.style.display = 'none'
    pauseBox.style.display = 'block'
    countdown.style.display = 'block'
  }
  if (event.key === 'r' && gameOver) {
    writeScores()
    gameOver = false
    isPaused = false
    mapp.p = false
    started = true
    gameOverBox.style.display = 'none'
    GameContainer.style.display = 'block'
    tutorialBox.style.display = 'block'
    pauseBox.style.display = 'none'
    countdown.style.display = 'block'
    win = false

    createBricks()
    bricks = [...document.querySelectorAll('.brick')]
    lives = 3
    stopTimer()
    countdownInterval = null
    currentTime = targetTime + 1
  } else if (event.key === 'r' && paused) {
    paused = false
    lives = 3
    pauseGame.style.display = 'none'
    started = false
    writeScores()
    brickTracker = 49
    currentTime = targetTime + 1
    ball.style.removeProperty('top')
    mapp.s = false
    stopTimer()
    countdownInterval = null
    tutorialBox.style.display = 'block'
    pauseBox.style.display = 'none'
    livesText.style.display = 'none'
    scoreText.style.display = 'none'
    createBricks()
    bricks = [...document.querySelectorAll('.brick')]
  } else if (event.key === 'r' && win) {
    lives = 3
    started = false
    win = false
    writeScores()
    brickTracker = 49
    currentTime = targetTime + 1
    ball.style.removeProperty('top')
    mapp.s = false
    // stopTimer()
    clearInterval(co)
    countdownInterval = null
    tutorialBox.style.display = 'block'
    pauseBox.style.display = 'none'
    livesText.style.display = 'none'
    scoreText.style.display = 'none'
    winBox.style.display = 'none'
    createBricks()
    bricks = [...document.querySelectorAll('.brick')]
  }

  document.addEventListener('keyup', event => {
    if (mapp.hasOwnProperty(event.key) && event.key !== 's') {
      mapp[event.key] = false
    }
  })
})

function MovePlayer () {
  if (mapp.ArrowLeft) {
    if (padelRect().x >= gameRect().x + padelSpeed) {
      padelPosX -= padelSpeed
    }
  }
  if (mapp.ArrowRight) {
    if (
      padelRect().x <
      gameRect().x + gameRect().width - padelRect().width - padelSpeed
    ) {
      padelPosX += padelSpeed
    }
  }
  padel.style.left = padelPosX + 'px'
}

function MoveBall () {
  if (posY >= GameContainer.offsetHeight - ball.offsetHeight) {
    mapp.s = false
    lives--
    if (lives < 1) {
      writeScores()
      gameOver = true
      started = false
      brickTracker = 49
      currentTime = targetTime + 1
      clearInterval(countdownInterval)
      countdownInterval = null
    }
    ball.style.removeProperty('top')
    return
  }

  let bounce = false
  let padelBounce = false

  if (DetectCollisionPadel(ball, padel)) {
    bounce = false
    console.log('Padel Collision')
    console.log(ballRect().top)
    console.log(padelRect().top)
    speedY *= -1
    if (!padelBounce) {
      console.log(speedX)
      console.log(ballRect().x)
      console.log(padelRect().x)

      padelBounce = true
      if (ballRect().x - 25 - padelRect().x < 0) {
        console.log('ACTIVE')
        if (speedX > 0) {
          speedX *= -1
        }
      } else if (ballRect().x - 25 - padelRect().x >= 0) {
        if (speedX < 0) {
          speedX *= -1
        }
      }
      posX += speedX
      posY += speedY
    }
  }

  let toRemove = []
  bricks.forEach(thing => {
    if (DetectCollision(ball, thing) && DetectCollisionSide(ball, thing)) {
      let thingRect = thing.getBoundingClientRect()
      toRemove.push(thing)
      padelBounce = false
      bounce = true
      speedX *= -1
      posX += speedX
      posY += speedY
    } else if (DetectCollision(ball, thing)) {
      padelBounce = false
      toRemove.push(thing)
      if (!bounce) {
        speedY *= -1
      }
      bounce = true
      posX += speedX
      posY += speedY
    }
  })

  if (toRemove.length !== 0) {
    toRemove.forEach(element => {
      idx = bricks.indexOf(element)
      bricks.splice(idx, 1)
      element.style.backgroundColor = 'black'
      brickTracker--
    })
  }

  posX += speedX
  posY += speedY

  if (posX < -1) {
    posX = 10
  } else if (posX > 344) {
    console.log('active')
    posX = 330
  }

  if (posX <= 0 || posX >= GameContainer.offsetWidth - ball.offsetWidth) {
    console.log('collide X')
    console.log(posX)

    speedX *= -1
  }

  if (posY <= 0 || posY >= GameContainer.offsetHeight - ball.offsetHeight) {
    console.log('collide Y')

    speedY *= -1
  }

  ball.style.left = posX + 'px'
  ball.style.top = posY + 'px'
}

function DetectCollisionPadel (elem1, elem2) {
  let rect1 = elem1.getBoundingClientRect()
  let rect2 = elem2.getBoundingClientRect()

  return (
    rect1.top + ball.offsetHeight >= rect2.top + padel.offsetHeight / 2 &&
    rect1.right <= rect2.right + ball.offsetWidth &&
    rect1.left + ball.offsetWidth >= rect2.left
  )
}

function DetectCollision (elem1, elem2) {
  let rect1 = elem1.getBoundingClientRect()
  let rect2 = elem2.getBoundingClientRect()
  return (
    rect1.top <= rect2.bottom &&
    rect1.right >= rect2.left &&
    rect1.bottom >= rect2.top &&
    rect1.left <= rect2.right
  )
}

function DetectCollisionSide (elem1, elem2) {
  let rect1 = elem1.getBoundingClientRect()
  let rect2 = elem2.getBoundingClientRect()
  const rectWidth = 20
  const rectHeight = 20

  const leftCollision =
    rect1.right >= rect2.left &&
    rect1.left < rect2.left &&
    rect1.bottom > rect2.top + rectHeight / 2 &&
    rect1.top < rect2.bottom - rectHeight / 2

  const rightCollision =
    rect1.left <= rect2.right &&
    rect1.right > rect2.right &&
    rect1.bottom > rect2.top + rectHeight / 2 &&
    rect1.top < rect2.bottom - rectHeight / 2

  const orCondition = leftCollision || rightCollision
  return orCondition
}
const offset = padelRect().left

function preGame () {
  ball.style.bottom = padel.style.bottom + padel.offsetHeight + 'px'
  ball.style.left =
    padel.offsetWidth / 2 -
    ball.offsetWidth / 2 +
    (padelRect().left - offset) +
    'px'
  padel.style.left + (padel.offsetWidth / 2 - ball.offsetWidth / 2) + 'px'
  posY = gameRect().height - ball.offsetHeight * 2
  posX = padel.offsetWidth / 2 + (padelRect().left - offset)
  countdown.style.display = 'none'
}

function GameEngine () {
  if (brickTracker <= 0) {
    writeScores()
    winBox.style.display = 'block'
    win = true
  }
  if (!paused && !win) {
    MovePlayer()
    if (mapp.s && !gameOver) {
      MoveBall()
      if (!started) {
        console.log('reached')
        if (!countdownInterval) {
          startTimer()
          console.log('YEEEW')
        }
        started = true
      }
      livesText.style.display = 'block'
      scoreText.style.display = 'block'
      livesText.innerText = `${lives} Lives`
    } else if (!gameOver) {
      preGame()
    } else {
      pauseTimer()
      console.log(score)
      countdown.style.display = 'none'
      gameOverBox.style.display = 'block'
      GameContainer.style.display = 'none'
      livesText.style.display = 'none'
      scoreText.style.display = 'none'
    }
    updateScore()
  }
  requestAnimationFrame(GameEngine)
}
requestAnimationFrame(GameEngine)

function colorBricksByRow () {
  const bricks = document.querySelectorAll('.brick')
  const colors = ['red', 'green', 'blue', 'yellow', 'purple', 'orange', 'pink']
  const bricksPerRow = 7
  bricks.forEach((brick, index) => {
    let rowIndex = Math.floor(index / bricksPerRow)
    brick.style.backgroundColor = colors[rowIndex % colors.length]
  })
}

function createBricks () {
  bricksContainer.innerHTML = ''
  for (let i = 0; i < brickTracker; i++) {
    bricks = undefined
    let brick = document.createElement('div')
    brick.setAttribute('class', 'brick')
    brick.setAttribute('id', i + 1)
    bricksContainer.appendChild(brick)
  }
  colorBricksByRow()
  bricks = [...document.querySelectorAll('.brick')]
}

function updateScore () {
  score = (49 - brickTracker) * 20
  scoreText.innerText = `Score:  ${score}`
}

function writeScores () {
  updateScore()
  scoreDisplay.forEach(thing => {
    thing.innerText = `Your Score Was ${score}`
  })
}

function stopTimer () {
  clearInterval(countdownInterval)
}

function timeOver () {
  gameOver = true
  started = false
  writeScores()
  brickTracker = 49
  currentTime = targetTime + 1
  ball.style.removeProperty('top')
  mapp.s = false
  stopTimer()
  countdownInterval = null
}
