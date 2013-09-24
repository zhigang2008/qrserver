(function($) {
$.extend($.validator.prototype, {
    showLabel: function(element, message) {
    }
});
$.extend($.validator.defaults, {
    errorClass: 'has-error',
    validClass: 'has-success',
    errorElement: 'span',
    highlight: function (element, errorClass, validClass) {
        var $element;
        if (element.type === 'radio') {
            $element = this.findByName(element.name);
        } else {
            $element = $(element);
        }
        $element.addClass(errorClass).removeClass(validClass);
        $element.parents("div.form-group").addClass("has-error");
    },
    unhighlight: function (element, errorClass, validClass) {
        var $element;
        if (element.type === 'radio') {
            $element = this.findByName(element.name);
        } else {
            $element = $(element);
        }
        $element.removeClass(errorClass).addClass(validClass);
        $element.parents("div.form-group").removeClass("has-error");
    },
    showErrors: function (errorMap, errorList) {
        $.each(this.successList, function (index, value) {
            $(value).popover('hide');
        });
        $.each(errorList, function (index, value) {
            var pop = $(value.element).popover({
                trigger: 'manual',
                content: value.message
            });
            pop.data('bs.popover').options.content = value.message;
            $(value.element).popover('show');
        });
        this.defaultShowErrors();
    }
});
}(jQuery));