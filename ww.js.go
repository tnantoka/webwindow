package webwindow

const WWJS = `(function() {
  document.addEventListener('click', function(e) {
    var elm = e.target;
    if (elm.tagName == 'A' && elm.target == '_blank') {
      e.preventDefault();
      var xhr = new XMLHttpRequest();
      xhr.open('POST', '/ww/open');
      xhr.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
      xhr.send('href=' + encodeURIComponent(elm.href));
    }
  });
})();`
