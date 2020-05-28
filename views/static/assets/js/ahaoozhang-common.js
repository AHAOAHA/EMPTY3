function openDownloadDialog(url, saveName) {
    if (typeof url == 'object' && url instanceof Blob) {
        url = URL.createObjectURL(url); // 创建blob地址
    }
    var aLink = document.createElement('a');
    aLink.href = url;
    aLink.download = saveName || ''; // HTML5新增的属性，指定保存文件名，可以不要后缀，注意，file:///模式下不会生效
    var event;
    if (window.MouseEvent) event = new MouseEvent('click');
    else {
        event = document.createEvent('MouseEvents');
        event.initMouseEvent('click', true, false, window, 0, 0, 0, 0, 0, false, false, false, false, 0, null);
    }
    aLink.dispatchEvent(event);
}

function csv2sheet(csv) {
    var sheet = {}; // 将要生成的sheet
    csv = csv.split('\n');
    csv.forEach(function (row, i) {
        row = row.split(',');
        if (i == 0) sheet['!ref'] = 'A1:' + String.fromCharCode(65 + row.length - 1) + (csv.length - 1);
        row.forEach(function (col, j) {
            sheet[String.fromCharCode(65 + j) + (i + 1)] = { v: col };
        });
    });
    return sheet;
}

// 将一个sheet转成最终的excel文件的blob对象，然后利用URL.createObjectURL下载
function sheet2blob(sheet, sheetName) {
    sheetName = sheetName || 'sheet1';
    var workbook = {
        SheetNames: [sheetName],
        Sheets: {}
    };
    workbook.Sheets[sheetName] = sheet;
    // 生成excel的配置项
    var wopts = {
        bookType: 'xlsx', // 要生成的文件类型
        bookSST: false, // 是否生成Shared String Table，官方解释是，如果开启生成速度会下降，但在低版本IOS设备上有更好的兼容性
        type: 'binary'
    };
    var wbout = XLSX.write(workbook, wopts);
    var blob = new Blob([s2ab(wbout)], { type: "application/octet-stream" });
    // 字符串转ArrayBuffer
    function s2ab(s) {
        var buf = new ArrayBuffer(s.length);
        var view = new Uint8Array(buf);
        for (var i = 0; i != s.length; ++i) view[i] = s.charCodeAt(i) & 0xFF;
        return buf;
    }
    return blob;
}

function make_college_options() {
    $.ajax({
        type: "GET",
        url: "/get_all_college_name?major_uid=" + document.getElementById("major_uid").value,
        dataType: "json",
        success: function (data) {
            var ss = document.getElementById("college_uid");
            $.each(data, function (index, val) {
                var op = document.createElement("option");
                op.setAttribute("value", val["CollegeUid"]);
                op.innerHTML = val["CollegeName"];
                ss.appendChild(op);
            })
        }
    })
}

function make_major_options() {
    $.ajax({
        type: "GET",
        url: "/get_all_major_name?college_uid=" + document.getElementById("college_uid").value,
        dataType: "json",
        success: function (data) {
            $("#major_uid").empty();
            var ss = document.getElementById("major_uid");
            var op = document.createElement("option");
            op.setAttribute("value", "");
            op.innerHTML = "请选择所属专业";
            ss.appendChild(op);
            
            $.each(data, function (index, val) {
                var op = document.createElement("option");
                op.setAttribute("value", val["MajorUid"]);
                op.innerHTML = val["MajorName"];
                ss.appendChild(op);
            })
        }
    })
}

function make_class_options() {
    $.ajax({
        type: "GET",
        url: "/get_all_class_name",
        dataType: "json",
        success: function (data) {
            $("#class_uid").empty();
            var ss = document.getElementById("class_uid");
            var op = document.createElement("option");
            op.setAttribute("value", "");
            op.innerHTML = "请选择所属班级";
            ss.appendChild(op);

            $.each(data, function (index, val) {
                var op = document.createElement("option");
                op.setAttribute("value", val["ClassUid"]);
                op.innerHTML = val["ClassName"];
                ss.appendChild(op);
            })
        }
    })
}

function make_teacher_options() {
    $.ajax({
        type: "GET",
        url: "/get_teacher_info?college_uid=" + document.getElementById("college_uid").value,
        dataType: "json",
        success: function (data) {
            $("#teacher_uid").empty();
            var ss = document.getElementById("teacher_uid");
            var op = document.createElement("option");
            op.setAttribute("value", "");
            op.innerHTML = "请选择授课教师";
            ss.appendChild(op);
            
            $.each(data, function (index, val) {
                var op = document.createElement("option");
                op.setAttribute("value", val["TeacherUid"]);
                op.innerHTML = val["TeacherName"];
                ss.appendChild(op);
            })
        }
    })
}

function make_course_options() {
    $.ajax({
        type: "GET",
        url: "/get_course_info?college_uid=" + document.getElementById("college_uid").value,
        dataType: "json",
        success: function (data) {
            $("#course_uid").empty();
            var ss = document.getElementById("course_uid");
            var op = document.createElement("option");
            op.setAttribute("value", "");
            op.innerHTML = "请选择课程";
            ss.appendChild(op);
            
            $.each(data, function (index, val) {
                var op = document.createElement("option");
                op.setAttribute("value", val["CourseUid"]);
                op.innerHTML = val["CourseName"];
                ss.appendChild(op);
            })
        }
    })
}

function make_major_and_teacher_and_course_options() {
    make_course_options();
    make_teacher_options();
    make_major_options();
}