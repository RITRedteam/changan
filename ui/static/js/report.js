function test() {
  var device = getParameterByName("device");

  var sel = document.getElementById('devices');
  var opts = sel.options;
  for (var opt, j = 0; opt = opts[j]; j++) {
    if (opt.value == device) {
      sel.selectedIndex = j;
      break;
    }
  }
}
