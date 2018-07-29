

// var setMenu = function (val) {
//     if(val === 1){
//         sessionStorage.setItem("msel", "1");
//     }else if(val === 2){
//         sessionStorage.setItem("msel", "2");
//     }else if(val === 3){
//         sessionStorage.setItem("msel", "3");
//     }else if(val === 4){
//         sessionStorage.setItem("msel", "4");
//     }

// }


var setMenu = function (val) {
    //alert(val)
    // var msel = sessionStorage.getItem("msel");
    if (val === 1) {
        document.getElementById("m1").classList.add('active');
        document.getElementById("m2").classList.remove('active');
        document.getElementById("m3").classList.remove('active');
        document.getElementById("m4").classList.remove('active');
    } else if (val === 2) {
        document.getElementById("m1").classList.remove('active');
        document.getElementById("m2").classList.add('active');
        //document.getElementById("m2").className += " active"
        document.getElementById("m3").classList.remove('active');
        document.getElementById("m4").classList.remove('active');

    } else if (val === 3) {
        document.getElementById("m1").classList.remove('active');
        document.getElementById("m2").classList.remove('active');
        document.getElementById("m3").classList.add('active');
        document.getElementById("m4").classList.remove('active');
    } else if (val === 4) {
        document.getElementById("m1").classList.remove('active');
        document.getElementById("m2").classList.remove('active');
        document.getElementById("m3").classList.remove('active');
        document.getElementById("m4").classList.add('active');
    }
    //  else{
    //     document.getElementById("m1").classList.add('active');
    //     document.getElementById("m2").classList.remove('active');
    //     document.getElementById("m3").classList.remove('active');
    //     document.getElementById("m4").classList.remove('active');
    // }
}

// var checkMenu = function(){
//     //alert(val)
//     var msel = sessionStorage.getItem("msel");
//     if(msel === "1"){
//         document.getElementById("m1").classList.add('active');
//         document.getElementById("m2").classList.remove('active');
//         document.getElementById("m3").classList.remove('active');
//         document.getElementById("m4").classList.remove('active');
//     }else if(msel === "2"){
//         document.getElementById("m1").classList.remove('active');
//         document.getElementById("m2").classList.add('active');
//         //document.getElementById("m2").className += " active"
//         document.getElementById("m3").classList.remove('active');
//         document.getElementById("m4").classList.remove('active');

//     }else if(msel === "3"){
//         document.getElementById("m1").classList.remove('active');
//         document.getElementById("m2").classList.remove('active');
//         document.getElementById("m3").classList.add('active');
//         document.getElementById("m4").classList.remove('active');
//     }else if(msel === "4"){
//         document.getElementById("m1").classList.remove('active');
//         document.getElementById("m2").classList.remove('active');
//         document.getElementById("m3").classList.remove('active');
//         document.getElementById("m4").classList.add('active');
//     }else{
//         document.getElementById("m1").classList.add('active');
//         document.getElementById("m2").classList.remove('active');
//         document.getElementById("m3").classList.remove('active');
//         document.getElementById("m4").classList.remove('active');
//     }
// }