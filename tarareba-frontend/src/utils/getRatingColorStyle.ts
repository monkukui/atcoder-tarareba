const getRatingColorStyle = (rating: number) => {
  let color
  if (rating >= 2800) {
    color = 'rgb(235, 51, 35)'
  } else if (rating >= 2400) {
    color = 'rgb(240, 135, 41)'
  } else if (rating >= 2000) {
    color = 'rgb(193, 191, 61)'
  } else if (rating >= 1600) {
    color = 'rgb(0, 32, 254)'
  } else if (rating >= 1200) {
    color = 'rgb(84, 189, 191)'
  } else if (rating >= 800) {
    color = 'rgb(55, 125, 34)'
  } else if (rating >= 400) {
    color = 'rgb(120, 67, 21)'
  } else {
    color = 'rgb(128, 128, 128)'
  }

  return { color: color, fontWeight: 'bold' }
}

export default getRatingColorStyle
