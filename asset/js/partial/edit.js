var editView = Backbone.View.extend({
    template: _.template($("#templateAdd").html()),
    render: function(id) {
        var $self = this
        $self.$el.html($self.template)
        showOneData(id)
        return $self
    },
    events:{
        "click #submitAdd": "doEdit"
    },
    doEdit: function() {
        var id = $("#id").val()
        var name = $("#name").val()
        var age = $("#age").val()
        var grade = $("#grade").val()
        var data  = {
            "_id": id,
            "name": name,
            "age": age,
            "grade": grade
        }
        if (id == "" || name == "" || age == "" || grade == "") {
            alert("Field can't empty")
        }else {
            $.ajax({
                url: "/api/updates",
                data: JSON.stringify(data),
                type: "PUT",
                dataType: "json",
                contentType: "application/json",
                success: function(v) {
                    alert("Update data is success")
                    app.navigate("", {trigger: true})
                },
                error: function(v, i){
                    console.log(v, i)
                }
            })
        }
    }
})