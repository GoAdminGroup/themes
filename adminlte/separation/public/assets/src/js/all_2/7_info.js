// ============================
// toastr initialization
// ============================

toastr.options = {
    closeButton: true,
    progressBar: true,
    showMethod: 'slideDown',
    timeOut: 4000
};

// ============================
// NProgress initialization
// ============================

NProgress.configure({parent: '#pjax-container'});

// ============================
// pjax initialization
// ============================

$.pjax.defaults.timeout = 5000;
$.pjax.defaults.maxCacheLength = 0;

$(document).pjax('a:not(a[target="_blank"])', {
    container: '#pjax-container'
});

$(document).on("pjax:click", "a.no-pjax", false);

$(document).on('pjax:timeout', function (event) {
    event.preventDefault();
});

$(document).on('submit', 'form[pjax-container]', function (event) {
    $.pjax.submit(event, '#pjax-container');
});

$(document).on('pjax:popstate', function () {

    $(document).one('pjax:end', function (event) {
        $(event.target).find('script[data-exec-on-popstate]').each(function () {
            $.globalEval(this.text || this.textContent || this.innerHTML || '');
        });
    });
});

$(document).on('pjax:send', function (xhr) {
    if (xhr.relatedTarget && xhr.relatedTarget.tagName && xhr.relatedTarget.tagName.toLowerCase() === 'form') {
        let submitBtn = $('form[pjax-container] :submit');
        if (submitBtn) {
            submitBtn.button('loading');
        }
    }
    NProgress.start();
});

$(document).on('pjax:complete', function (xhr) {
    if (xhr.relatedTarget && xhr.relatedTarget.tagName && xhr.relatedTarget.tagName.toLowerCase() === 'form') {
        let submitBtn = $('form[pjax-container] :submit');
        if (submitBtn) {
            submitBtn.button('reset');
        }
    }
    NProgress.done();
    updateNavURL();
});

// ============================
// top nav bar buttons
// ============================

let fullpageBtn = $('.fullpage-btn');
let exitFullpageBtn = $('.exit-fullpage-btn');

fullpageBtn.on('click', function () {
    launchFullscreen(document.documentElement);
    fullpageBtn.hide();
    exitFullpageBtn.show();
});

exitFullpageBtn.on('click', function () {
    exitFullscreen();
    exitFullpageBtn.hide();
    fullpageBtn.show();
});

function launchFullscreen(element) {
    if (element.requestFullscreen) {
        element.requestFullscreen();
    } else if (element.mozRequestFullScreen) {
        element.mozRequestFullScreen();
    } else if (element.msRequestFullscreen) {
        element.msRequestFullscreen();
    } else if (element.webkitRequestFullscreen) {
        element.webkitRequestFullScreen();
    }
}

function exitFullscreen() {
    if (document.exitFullscreen) {
        document.exitFullscreen();
    } else if (document.msExitFullscreen) {
        document.msExitFullscreen();
    } else if (document.mozCancelFullScreen) {
        document.mozCancelFullScreen();
    } else if (document.webkitExitFullscreen) {
        document.webkitExitFullscreen();
    }
}

$('.container-refresh').on('click', function () {
    $.pjax.reload('#pjax-container');
    toastr.success(toastMsg);
});

// ============================
// sidebar initialization
// ============================

let sidebarMenuA = $('.sidebar-menu a');

$(function () {
    $('.sidebar-menu li:not(.treeview) > a').on('click', function () {
        let parent = $(this).parent().addClass('active');
        parent.siblings('.treeview.active').find('> a').trigger('click');
        parent.siblings().removeClass('active').find('li').removeClass('active');
    });

    $('[data-toggle="popover"]').popover();

    addOrRemoveLeftRightNavBtn(false);

    initMaxNavWrapperWidth();
});

$(window).resize(function() {
    initMaxNavWrapperWidth();
    addOrRemoveLeftRightNavBtn(!checkNavLength());
});

sidebarMenuA.on('click', function () {

    let link = $(this).attr('href');
    if (link !== '#' && link.indexOf('http') === -1 && !checkNavExist(link)) {

        addOrRemoveLeftRightNavBtn(!checkNavLength());

        removeActive();

        addNavTab(link, $(this).html().replace('<i/><span>', '<i/>&nbsp&nbsp&nbsp<span>'))

        moveToRight();
    }
});

$('a.new-tab-link').on('click', function () {
    listenerForAddNavTab($(this).attr('href'), $(this).attr('data-title'))
});

$('.navbar-nav-btn-left').on('click', function () {
    moveToLeft();
});

function moveToLeft() {
    let nav = $(".nav.nav-tabs.nav-addtabs");
    let marginLeft = parseInt(nav.css("marginLeft"));
    if (marginLeft < 0) {
        if (marginLeft < -50 && marginLeft > -100) {
            nav.css("marginLeft", "0px");
        } else {
            nav.css("marginLeft", (marginLeft + 50) + "px");
        }
    }
}

$('.navbar-nav-btn-right').on('click', function () {
    moveToRight();
});

function moveToRight() {
    let nav = $(".nav.nav-tabs.nav-addtabs");
    let marginLeft = parseInt(nav.css("margin-left"));
    let ulWidth = getNavULwidth();
    let padding = ulWidth - maxNavWrapperWidth;
    if (padding > 0 && marginLeft + padding != 0) {
        if (padding + marginLeft < 100) {
            nav.css("margin-left", -padding + "px");
        } else {
            nav.css("margin-left", (marginLeft - 50) + "px");
        }
    }
}

let showNav = false;

function addOrRemoveLeftRightNavBtn(add) {
    if (add) {
        if (!showNav) {
            $('.navbar-nav-btn-right').show();
            $('.navbar-nav-btn-left').show();
            showNav = true;  
        }
    } else {
        if (showNav) {
            $('.navbar-nav-btn-right').hide();
            $('.navbar-nav-btn-left').hide();
            $(".nav.nav-tabs.nav-addtabs").css("margin-left", "0px");          
            showNav = false;
        }        
    }
}

function getNavULwidth() {
    let lis = $(".nav.nav-tabs.nav-addtabs li");
    let width = 0;
    for (let i = 0; i < lis.length; i++) {
        width += $(lis[i]).width()
    }
    return width
}

function listenerForAddNavTab(link, content) {
    if (link !== '#' && link.indexOf('http') === -1 && !checkNavExist(link)) {

        addOrRemoveLeftRightNavBtn(!checkNavLength());

        removeActive();

        if (content === "") {

            let sidebarMenus = sidebarMenuA;
            let re = new RegExp("\\?(.*)");

            for (let i = 0; i < sidebarMenus.length; i++) {
                if (link.replace(re, '') === $(sidebarMenus[i]).attr('href')) {
                    content = $(sidebarMenus[i]).html().replace('<i/><span>', '<i/>&nbsp&nbsp&nbsp<span>');
                    break
                }
            }
        }

        if (content !== "") {
            addNavTab(link, content)
            moveToRight();
        }
    }
}

function addNavTab(link, content) {
    let addElement = $('<li class="active">\n' +
        '<a href="' + link + '">\n' +
        '<span>' + content + '</span>\n' +
        '</a><i class="close-tab fa fa-remove"></i>\n' +
        '</li>');

    addElement.find('.close-tab').on('click', function () {
        let li = $(this).parent();
        if (li.hasClass('active')) {
            if (li.prev().length > 0) {
                li.prev().addClass('active');
                $.pjax({url: li.prev().find('a').attr('href'), container: '#pjax-container'});
            } else if (li.next().length > 0) {
                li.next().addClass('active');
                $.pjax({url: li.next().find('a').attr('href'), container: '#pjax-container'});
            }
        }
        li.remove();
        addOrRemoveLeftRightNavBtn(!checkNavLength());
        moveToLeft();
    });
    addElement.on('mouseover', function () {
        if ($(this).children('i')) {
            $(this).children('i').show();
        }
    });
    addElement.on('mouseout', function () {
        if ($(this).children('i')) {
            $(this).children('i').hide();
        }
    });
    addElement.on('click', function () {
        removeActive();
        $(this).addClass('active');
    });

    addElement.appendTo('.nav-addtabs');
}

function checkNavExist(link) {
    let navs = $('.nav-addtabs li');
    for (let i = 0; i < navs.length; i++) {
        if (parseURL($(navs[i]).find('a').attr('href')) === link.split("?")[0]) {
            removeActive();
            $(navs[i]).addClass('active');
            return true;
        }
    }
    return false;
}

function parseURL(url) {
    let t = url.substring(url.indexOf("//") + 2);
    return t.substring(t.indexOf("/")).split("?")[0];
}

function updateNavURL() {
    let navs = $('.nav-addtabs li');
    for (let i = 0; i < navs.length; i++) {
        if ($(navs[i]).hasClass('active')) {
            $(navs[i]).find('a').attr('href', location.href);
        }
    }
}

function removeActive() {
    let lis = $('.nav-addtabs li');
    for (let i = 0; i < lis.length; i++) {
        $(lis[i]).removeClass('active');
    }
}

let maxNavWrapperWidth = 0;

function initMaxNavWrapperWidth() {
    let navWidth = $('#firstnav').width();
    let menuWidth = $('.navbar-custom-menu').width();
    maxNavWrapperWidth = (navWidth - menuWidth) * 0.7
    $('.nav-tabs-content').css("max-width", maxNavWrapperWidth + "px");
    $('.nav.nav-tabs.nav-addtabs').css("width", maxNavWrapperWidth + 800 + "px");
    return maxNavWrapperWidth
}

function checkNavLength() {
    let ulWidth = getNavULwidth();
    return ulWidth + 50 < maxNavWrapperWidth;
}

const fixedKey = "go_admin__sidebar_fixed";

$(function () {
    let isFixed = window.localStorage.getItem(fixedKey);
    if (isFixed === "true") {
        $('.main-sidebar').css('position', 'fixed');
        $('.main-header .logo').css('position', 'fixed');
        $('.fixed-btn').attr('data-click', 'true');
    }
});


$('.fixed-btn').on('click', function () {
    let clicked = $(this).attr('data-click');
    if (clicked === "false") {
        $('.main-sidebar').css('position', 'fixed');
        $('.main-header .logo').css('position', 'fixed');
        $(this).attr('data-click', 'true');
        window.localStorage.setItem(fixedKey, "true");
        $(this).css('background-color', '#f3f3f3');
    } else {
        $('.main-sidebar').css('position', '');
        $('.main-header .logo').css('position', '');
        $(this).attr('data-click', 'false');
        window.localStorage.removeItem(fixedKey);
        $(this).css('background-color', 'white');
    }
});