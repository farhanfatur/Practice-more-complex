var addView = Backbone.View.extend({
    events:  {
        "click #submitAdd": "doAdd"
    },
    template: _.template($("#templateAdd").html()),
    render: function() {
        var $self = this
        $self.$el.html($self.template)
        return $self
    },
    doAdd: function(ev) {
        ev.preventDefault()
        var data = {
            "name": $("#name").val(),
            "age": parseInt($("#age").val()),
            "grade": parseInt($("#grade").val())
        }
        // var jsonNum = JSON.stringify(data)
        if (data == "") {
            alert("Field is empty")
        }else {
            // console.log(jsonNum)
            console.log(data)
            $.ajax({
                url: "/api/inserts",
                type: "POST",
                data: JSON.stringify(data),
                dataType: "json",
                contentType: "application/json",
                success: function(v, i) {
                    alert("Insert is success")
                    showData()
                    app.navigate("", {trigger: true})
                },
                error: function(v, i) {
                    console.log(v, i)
                }
            })
        }
    }
})