var map;
var currentMarker;

google.load('visualization', '1.0', {'packages':['corechart']});

$(document).ready(function() {
  var mapOptions = {
    zoom: 5,
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
    var state = $("<span>", {class: 'name'});
    state.html(response.state);
    infobox.append(state);
    var header = $("<h1>");
    header.html("Tract " + response.tract + ", " + response.county + " County");
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
  wrapper.append('<span class="name">' + report.name + '</span>');
  var data = new google.visualization.DataTable();
  data.addColumn('string', 'Race');
  data.addColumn('number', 'Count');
  $.each(report.parts, function(key, variable) {
    if (parseInt(variable) > 10) {
      data.addRow([key, parseInt(variable)]);
    }
  });
  data.sort({column: 1, desc: true});
  var chartDiv = $("<div>");
  chartDiv.addClass("chart");
  var chart = new google.visualization.BarChart(chartDiv[0]);
  wrapper.append(chartDiv);
  chart.draw(data, {
    width: 450,
    animation: {duration: 1},
    legend: {position: 'none'},
    chartArea: {left: 100, height: "80%"}
  });
  return wrapper;
}
