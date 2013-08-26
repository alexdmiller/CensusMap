var map;
var currentMarker;

$(document).ready(function() {
  var mapOptions = {
    zoom: 4,
    center: new google.maps.LatLng(46.619, -120),
    mapTypeId: google.maps.MapTypeId.ROADMAP
  };
  map = new google.maps.Map(document.getElementById('map-canvas'),
    mapOptions);
  google.maps.event.addListener(map, 'click', onMapClick);
  setMarker(46.619, -120);
});

function onMapClick(event) {
  setMarker(event.latLng.mb, event.latLng.nb);
}

function setMarker(longitude, latitude) {
   if (currentMarker) {
    currentMarker.setMap(null);
  }
  currentMarker = new google.maps.Marker({position: new google.maps.LatLng(longitude, latitude), map: map});
  updateInfoBox(longitude, latitude);
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
    infobox.html("");
    response = JSON.parse(response);
    var header = $("<h1>");
    header.html(response.county + ", " + response.state + " (Tract: " + response.tract + ")");
    infobox.append(header);
    console.log(response);
    $.each(response.reports, function(i, report) {
      switch (report.kind) {
        case 'plain_value':
        infobox.append(renderPlainValueReport(report));
        break;
        case 'composition':
        infobox.append(renderCompositionReport(report));
        break;
      }
    });
    infobox.removeClass('loading');
  });
}

function commaSeparateNumber(val) {
  while (/(\d+)(\d{3})/.test(val.toString())) {
    val = val.toString().replace(/(\d+)(\d{3})/, '$1'+','+'$2');
  }
  return val;
}

function renderPlainValueReport(report) {
  var wrapper = $("<div>");
  wrapper.addClass("report");
  $.each(report.vars, function(key, variable) {
    variableWrapper = $("<div>");
    variableWrapper.addClass('single-value');
    variableWrapper.html('<span class="name">' + key + '</span><span class="value">' + commaSeparateNumber(variable) + '</span>');
    wrapper.append(variableWrapper);
  });
  console.log(wrapper);
  return wrapper;
}

function renderCompositionReport(report) {
  var wrapper = $("<div>");
  wrapper.addClass("report");
  $.each(report.parts, function(key, variable) {
    variableWrapper = $("<div>");
    variableWrapper.html(key + ": " + variable);
    wrapper.append(variableWrapper);
  });
  console.log(wrapper);
  return wrapper;
}
