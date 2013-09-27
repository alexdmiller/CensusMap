var map;
var currentMarker;

google.load('visualization', '1.0', {'packages':['corechart']});

$(document).ready(function() {
  // set up map
  var mapOptions = {
    zoom: 5,
    center: new google.maps.LatLng(39.50404070558415, -97.1630859375),
    mapTypeId: google.maps.MapTypeId.ROADMAP
  };
  map = new google.maps.Map(document.getElementById('map-canvas'),
    mapOptions);
  google.maps.event.addListener(map, 'click', onMapClick);
  
  $("#instructions").offset({
    top: $(window).height() / 2 - $("#instructions").height() / 2,
    left: $(window).width() / 2 - $("#instructions").width() / 2,
  });

  $("#instructions button").click(function() {
    $("#instructions").hide();
  });

  $("#info-box").hide();
  $("#info-box").height($(window).height() - 100);
  $(window).resize(function() {
    $("#info-box").height($(window).height() - 100);
  });
});

function onMapClick(event) {
  $("#instructions").hide();
  $("#info-box").show();
  setMarker(event.latLng.nb, event.latLng.ob);
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
  var currentScroll = infobox.scrollTop();
  var spinner = new Spinner().spin(infobox[0]);
  infobox.addClass('loading');
  $.ajax("api/census", {
    data: {
      lat: latitude,
      long: longitude
    }
  }).done(function(response) {
    infobox.html("");
    try {
      response = JSON.parse(response);
      var state = $("<span>", {class: 'name'});
      state.html(response.state);
      infobox.append(state);
      var header = $("<h1>");
      header.html("Tract " + response.tract + ", " + response.county + " County");
      infobox.append(header);
      $.each(response.reports, function(i, report) {
        switch (report.kind) {
          case 'plain_value':
          infobox.append(renderPlainValueReport(report));
          break;
          case 'composition':
          infobox.append(renderCompositionReport(report));
          break;
          case 'population_pyramid':
          infobox.append(renderPopulationPyramidReport(report));
          break;
        }
      });
      infobox.removeClass('loading');
      infobox.scrollTop(currentScroll);
    } catch (error) {
      var errorMessage = $("<div>");
      errorMessage.addClass('error');
      errorMessage.append("<h1>Error</h1>");
      errorMessage.append("<p>" + response + "</p>");
      infobox.append(errorMessage);
      infobox.removeClass('loading');
    }
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
  return wrapper;
}

function renderCompositionReport(report) {
  var wrapper = $("<div>");
  wrapper.addClass("report");
  wrapper.append('<span class="name">' + report.name + '</span>');
  var data = new google.visualization.DataTable();
  data.addColumn('string', 'Title');
  data.addColumn('number', 'Count');
  $.each(report.parts, function(key, variable) {
    if (!report.dropZeros || parseInt(variable[1]) > 0) {
      data.addRow([variable[0], parseInt(variable[1])]);
    }
  });
  if (report.sorted) {
    data.sort({column: 1, desc: true});  
  }
  var chartDiv = $("<div>");
  chartDiv.addClass("chart");
  var chart = new google.visualization.BarChart(chartDiv[0]);
  wrapper.append(chartDiv);
  if (report.display == "list") {
    chart.draw(data, {
      width: 450,
      height: 450,
      animation: {duration: 1},
      legend: {position: 'none'},
      chartArea: {width: "100%", height: "90%"},
      vAxis: {textPosition: 'in', textStyle: {fontSize: 15}},
      bar: {groupWidth: "90%"},
      colors: ['#B8E6E6']
    });
  } else {
    chart.draw(data, {
      width: 450,
      animation: {duration: 1},
      legend: {position: 'none'},
      chartArea: {left: 130, width: 300, height: "80%"}
    });
  }
  
  return wrapper;
}

function renderPopulationPyramidReport(report) {
  var wrapper = $("<div>");
  wrapper.addClass("report");
  wrapper.append('<span class="name">' + report.name + '</span>');
  var data = new google.visualization.DataTable();
  data.addColumn('string', 'Title');
  data.addColumn('number', 'Male');
  data.addColumn('number', 'Female');
  $.each(report.ages, function(key, variable) {
    data.addRow([variable[0], parseInt(variable[1]), -parseInt(variable[2])]);
  });
  var chartDiv = $("<div>");
  chartDiv.addClass("chart");
  var chart = new google.visualization.BarChart(chartDiv[0]);
  wrapper.append(chartDiv);
  chart.draw(data, {
    width: 450,
    height: 400,
    animation: {duration: 1},
    legend: {position: 'none'},
    chartArea: {left: 120, height: "100%"},
     isStacked: true,
    vAxis: {
      direction: -1
    }
  });
  return wrapper;
}
