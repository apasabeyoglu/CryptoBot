  $(function() {

    "use strict";
    $('select').material_select();

    // Commom, Translation & Horizontal Dropdown
    $('.dropdown-button, .translation-button, .dropdown-menu').dropdown({
      inDuration: 300,
      outDuration: 225,
      constrainWidth: false,
      hover: true,
      gutter: 0,
      belowOrigin: true,
      alignment: 'left',
      stopPropagation: false
    });
    // Notification, Profile & Settings Dropdown
    $('.notification-button, .profile-button, .dropdown-settings').dropdown({
      inDuration: 300,
      outDuration: 225,
      constrainWidth: false,
      hover: true,
      gutter: 0,
      belowOrigin: true,
      alignment: 'right',
      stopPropagation: false
    });

    // Materialize scrollSpy
    $('.scrollspy').scrollSpy();

    // Materialize tooltip
    $('.tooltipped').tooltip({
      delay: 50
    });

    //Main Left Sidebar Menu
    $('.sidebar-collapse').sideNav({
      edge: 'left', // Choose the horizontal origin
    });

    // Overlay Menu (Full screen menu)
    $('.menu-sidebar-collapse').sideNav({
      menuWidth: 240,
      edge: 'left', // Choose the horizontal origin
      //closeOnClick:true, // Set if default menu open is true
      menuOut: false // Set if default menu open is true
    });

    // Pikadate datepicker
    $('.datepicker').pickadate({
      selectMonths: true, // Creates a dropdown to control month
      selectYears: 15 // Creates a dropdown of 15 years to control year
    });

    // Perfect Scrollbar
    $('select').not('.disabled').material_select();
    var leftnav = $(".page-topbar").height();
    var leftnavHeight = window.innerHeight - leftnav;
    if (!$('#slide-out.leftside-navigation').hasClass('native-scroll')) {
      $('.leftside-navigation').perfectScrollbar({
        suppressScrollX: true
      });
    }

    // Fullscreen
    function toggleFullScreen() {
      if ((document.fullScreenElement && document.fullScreenElement !== null) ||
        (!document.mozFullScreen && !document.webkitIsFullScreen)) {
        if (document.documentElement.requestFullScreen) {
          document.documentElement.requestFullScreen();
        } else if (document.documentElement.mozRequestFullScreen) {
          document.documentElement.mozRequestFullScreen();
        } else if (document.documentElement.webkitRequestFullScreen) {
          document.documentElement.webkitRequestFullScreen(Element.ALLOW_KEYBOARD_INPUT);
        }
      } else {
        if (document.cancelFullScreen) {
          document.cancelFullScreen();
        } else if (document.mozCancelFullScreen) {
          document.mozCancelFullScreen();
        } else if (document.webkitCancelFullScreen) {
          document.webkitCancelFullScreen();
        }
      }
    }

    $('.toggle-fullscreen').click(function() {
      toggleFullScreen();
    });

    // Detect touch screen and enable scrollbar if necessary
    function is_touch_device() {
      try {
        document.createEvent("TouchEvent");
        return true;
      } catch (e) {
        return false;
      }
    }
    if (is_touch_device()) {
      $('#nav-mobile').css({
        overflow: 'auto'
      })
    }
  });