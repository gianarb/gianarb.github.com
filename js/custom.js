jQuery(document).ready(function () {

  var navigatorSwitchActive = function() {
    var els = document.querySelectorAll("#top a");
    $(".nav").find("#top .active").removeClass("active");
    for (var ii = 0; ii < els.length; ii++) {
      if (els[ii].href === location.href) {
        $(els[ii]).parent().addClass("active");
        return;
      }
    }
  };
  navigatorSwitchActive();
});
