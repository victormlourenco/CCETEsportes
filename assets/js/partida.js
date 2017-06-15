(function () {

    // Define a function that updates all relative dates defined by <time class='cw-relative-date'>
    var updateAllRelativeDates = function () {
        $('time').each(function (i, e) {
            if ($(e).attr("class") == 'cw-relative-date') {

                // Initialise momentjs
                moment.locale('pt-br');

                // Grab the datetime for the element and compare to now                    
                var time = moment($(e).attr('datetime')).add(3, 'hours');
                $(e).html('<span>' + time.format('LLL') + '</span>');
            }
        });
    };

    // Update all dates initially
    updateAllRelativeDates();

    // Register the timer to call it again every minute
    setInterval(updateAllRelativeDates, 60000);

})();