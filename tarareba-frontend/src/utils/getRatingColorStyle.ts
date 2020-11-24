const getRatingColorStyle = (rating: number) => {
  if (rating >= 2800) {
    return { backgroundColor: 'rgb(243, 177, 176)' }
  }
  if (rating >= 2400) {
    return { backgroundColor: 'rgb(250, 218, 183)' }
  }
  if (rating >= 2000) {
    return { backgroundColor: 'rgb(237, 235, 184)' }
  }
  if (rating >= 1600) {
    return { backgroundColor: 'rgb(177, 180, 250)' }
  }
  if (rating >= 1200) {
    return { backgroundColor: 'rgb(190, 235, 235)' }
  }
  if (rating >= 800) {
    return { backgroundColor: 'rgb(186, 216, 181)' }
  }
  if (rating >= 400) {
    return { backgroundColor: 'rgb(214, 198, 180)' }
  }
  return { backgroundColor: 'rgb(217, 217, 217)' }
}

export default getRatingColorStyle
