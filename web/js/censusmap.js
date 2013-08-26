var map;
var currentMarker;

$(document).ready(function() {
  var mapOptions = {
    zoom: 1,
    center: new google.maps.LatLng(46.619, -120),
    mapTypeId: google.maps.MapTypeId.ROADMAP
  };
  map = new google.maps.Map(document.getElementById('map-canvas'),
    mapOptions);
  google.maps.event.addListener(map, 'click', onMapClick);
  updateInfoBox(47.598755, -122.332764);
});

function onMapClick(event) {
  if (currentMarker) {
    currentMarker.setMap(null);
  }
  currentMarker = new google.maps.Marker({position: event.latLng, map: map});
  updateInfoBox(event.latLng.mb, event.latLng.nb);
}

function updateInfoBox(latitude, longitude) {
  var infobox = $("#info-box");
  var spinner = new Spinner().spin(infobox[0]);
  infobox.addClass('loading');
  $.ajax("api/census", {
    data: {
      lat: latitude,
      long: longitude
    }
  }).done(function(response) {
    console.log(response);
    infobox.html(response);
    infobox.removeClass('loading');
  });
}

