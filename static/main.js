const videoPlayer = document.getElementById('videoPlayer');

// Add event listener to handle errors
videoPlayer.addEventListener('error', () => {
    console.error('Video playback error');
});

// Add event listener to check when the video has ended
videoPlayer.addEventListener('ended', () => {
    console.log('Video ended');
});
