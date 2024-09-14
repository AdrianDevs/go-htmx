function navigateBack(parentPageURL) {
  console.group('navigateBack');
  console.log('parentPageURL', parentPageURL);
  console.log('referrer url', document.referrer);

  // Check if the referrer is the parent page
  // if (document.referrer === parentPageURL) {
  //     // If so then navigate back
  //     history.back();
  // } else {
  //     // If not, navigate to the parent page
  window.location.href = parentPageURL;
  // }
}

function updateTabTitle(title) {
  console.log('updateTabTitle', title);
  document.title = title;
}
