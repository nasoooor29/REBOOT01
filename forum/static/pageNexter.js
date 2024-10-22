let noPages = Number(document.getElementById('pages').innerText)
let currentpage = Number(document.getElementById('page-number').innerText)
let prevButton = document.getElementById('prev-page')
let nextButton = document.getElementById('next-page')
if (noPages === 1) {
  prevButton.style.opacity = 0.7
  nextButton.style.opacity = 0.7
  prevButton.disabled = true
  nextButton.disabled = true
} else if (currentpage === 1) {
  prevButton.style.opacity = 0.7
  prevButton.disabled = true
} else if (currentpage === noPages) {
  nextButton.style.opacity = 0.7
  nextButton.disabled = true
}

function next () {
  check = window.location.href
  added = ''
  if (check.length > 32) {
    if (currentpage >= 10) {
      added = check.slice(30)
    } else {
      added = check.slice(29)
    }
  }
  console.log(added)

  currentpage++
  window.location.href = `/posts/${currentpage}${added}`
}
function prev () {
  check = window.location.href
  added = ''
  if (check.length > 32) {
    if (currentpage >= 10) {
      added = check.slice(30)
    } else {
      added = check.slice(29)
    }
  }
  console.log(added)

  currentpage--
  window.location.href = `/posts/${currentpage}${added}`
}

nextButton.addEventListener('click', next)
prevButton.addEventListener('click', prev)
