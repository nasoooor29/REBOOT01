function likePost (postID, logged) {
  console.log(logged)

  if (logged === 'false') {
    window.location.href = '/signIn'
  }

  fetch(`/like/${postID}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then(res => res.json())
    .then(data => {
      likesTag = document.getElementById(`likes-${postID}`)
      dislikesTag = document.getElementById(`dislikes-${postID}`)

      likesTag.innerText = data.likes
      dislikesTag.innerText = data.dislikes
    })
    .catch(error => console.error('Error fetching data:', error))

  // noDislikes = Number(dislikesTag.innerText)
  // if (dislikesTag.dataset.disliked === 'active') {
  //   dislikesTag.innerText = --noDislikes
  //   dislikesTag.dataset.disliked = 'inactive'
  // }

  // noLikes = Number(likesTag.innerText)
  // if (likesTag.dataset.liked === 'active') {
  //   likesTag.dataset.liked = 'inactive'
  //   likesTag.innerText = --noLikes
  // } else {
  //   likesTag.dataset.liked = 'active'
  //   likesTag.innerText = ++noLikes
  // }
}

function dislikePost (postID, logged) {
 
  if (logged === 'false') {
    window.location.href = '/signIn'
  }
  fetch(`/dislike/${postID}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then(res => res.json())
    .then(data => {
      likesTag = document.getElementById(`likes-${postID}`)
      dislikesTag = document.getElementById(`dislikes-${postID}`)

      likesTag.innerText = data.likes
      dislikesTag.innerText = data.dislikes
    })
    .catch(error => console.error('Error fetching data:', error))
  // likesTag = document.getElementById('likes')
  // dislikesTag = document.getElementById('dislikes')

  // noLikes = Number(likesTag.innerText)
  // if (likesTag.dataset.liked === 'active') {
  //   likesTag.innerText = --noLikes
  //   likesTag.dataset.likes = 'inactive'
  // }

  // noDislikes = Number(dislikesTag.innerText)
  // if (dislikesTag.dataset.disliked === 'active') {
  //   dislikesTag.dataset.disliked = 'inactive'
  //   dislikesTag.innerText = --noDislikes
  // } else {
  //   dislikesTag.dataset.disliked = 'active'
  //   dislikesTag.innerText = ++noDislikes
  // }
}

function likeSinglePost (postID) {
  fetch(`/like/${postID}`, {
    method: 'POST'
  })
    .then(res => res.json())
    .then(data => {
      likesTag = document.getElementById(`count-like-${postID}`)
      dislikesTag = document.getElementById(`count-dislike-${postID}`)

      likesTag.innerText = data.likes
      dislikesTag.innerText = data.dislikes
    })
    .catch(error => console.error('Error fetching data:', error))
}

function dislikeSinglePost (postID) {
  fetch(`/dislike/${postID}`, {
    method: 'POST'
  })
    .then(res => res.json())
    .then(data => {
      likesTag = document.getElementById(`count-like-${postID}`)
      dislikesTag = document.getElementById(`count-dislike-${postID}`)

      likesTag.innerText = data.likes
      dislikesTag.innerText = data.dislikes
    })
    .catch(error => console.error('Error fetching data:', error))
}

function likeComment (commentID) {
  fetch(`/likeComment/${commentID}`, {
    method: 'POST'
  })
    .then(res => res.json())
    .then(data => {
      likesTag = document.getElementById(`comment-likes-${commentID}`)
      dislikesTag = document.getElementById(`comment-dislikes-${commentID}`)

      likesTag.innerText = data.likes
      dislikesTag.innerText = data.dislikes
    })
    .catch(error => console.error('Error fetching data:', error))
}

function dislikeComment (commentID) {
  fetch(`/dislikeComment/${commentID}`, {
    method: 'POST'
  })
    .then(res => res.json())
    .then(data => {
      likesTag = document.getElementById(`comment-likes-${commentID}`)
      dislikesTag = document.getElementById(`comment-dislikes-${commentID}`)

      likesTag.innerText = data.likes
      dislikesTag.innerText = data.dislikes
    })
    .catch(error => console.error('Error fetching data:', error))
}
