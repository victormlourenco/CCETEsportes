(function () {

    // Define a function that updates all relative dates defined by <time class='cw-relative-date'>
    var updateAllRelativeDates = function () {
        $('time').each(function (i, e) {
            if ($(e).attr("class") == 'cw-relative-date') {

                // Initialise momentjs
                var now = moment();
                moment.locale('pt-br');

                // Grab the datetime for the element and compare to now                    
                var time = moment($(e).attr('datetime'));
                var diff = now.diff(time, 'days');

                // If less than one day ago/away use relative, else use calendar display
                if (diff <= 1 && diff >= -1) {
                    $(e).html('<span>' + time.from(now) + '</span>');
                } else {
                    $(e).html('<span>' + time.calendar() + '</span>');
                }
            }
        });
    };

    // Update all dates initially
    updateAllRelativeDates();

    // Register the timer to call it again every minute
    setInterval(updateAllRelativeDates, 60000);

})();