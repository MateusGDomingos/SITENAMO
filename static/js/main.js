(function () {
    'use strict';

    var contador = document.querySelector('.contador-grid');
    if (contador) {
        var dataInicioStr = contador.getAttribute('data-inicio');
        var dataInicio = new Date(dataInicioStr);

        var campos = {
            anos: document.getElementById('anos'),
            meses: document.getElementById('meses'),
            dias: document.getElementById('dias'),
            horas: document.getElementById('horas'),
            minutos: document.getElementById('minutos'),
            segundos: document.getElementById('segundos')
        };

        function calcularDiferenca(inicio, agora) {
            var diffMs = agora - inicio;
            if (diffMs < 0) diffMs = 0;

            var segundosTotais = Math.floor(diffMs / 1000);
            var anos = Math.floor(segundosTotais / (365.25 * 24 * 3600));
            var resto = segundosTotais % Math.floor(365.25 * 24 * 3600);
            var meses = Math.floor(resto / (30.44 * 24 * 3600));
            resto = resto % Math.floor(30.44 * 24 * 3600);
            var dias = Math.floor(resto / (24 * 3600));
            resto = resto % (24 * 3600);
            var horas = Math.floor(resto / 3600);
            resto = resto % 3600;
            var minutos = Math.floor(resto / 60);
            var segundos = resto % 60;

            return { anos: anos, meses: meses, dias: dias, horas: horas, minutos: minutos, segundos: segundos };
        }

        function formatar(n) {
            return n < 10 ? '0' + n : String(n);
        }

        function atualizar() {
            var agora = new Date();
            var d = calcularDiferenca(dataInicio, agora);
            campos.anos.textContent = d.anos;
            campos.meses.textContent = formatar(d.meses);
            campos.dias.textContent = formatar(d.dias);
            campos.horas.textContent = formatar(d.horas);
            campos.minutos.textContent = formatar(d.minutos);
            campos.segundos.textContent = formatar(d.segundos);
        }

        atualizar();
        setInterval(atualizar, 1000);
    }

    var lightbox = document.getElementById('lightbox');
    var lightboxImg = lightbox ? lightbox.querySelector('.lightbox-img') : null;
    var galeriaItens = document.querySelectorAll('.galeria-item');
    var indiceAtual = 0;

    function abrirLightbox(indice) {
        if (!lightbox || galeriaItens.length === 0) return;
        indiceAtual = indice;
        var item = galeriaItens[indice];
        lightboxImg.src = item.getAttribute('data-src');
        lightboxImg.alt = item.getAttribute('data-alt');
        lightbox.classList.add('active');
        lightbox.setAttribute('aria-hidden', 'false');
        document.body.style.overflow = 'hidden';
    }

    function fecharLightbox() {
        if (!lightbox) return;
        lightbox.classList.remove('active');
        lightbox.setAttribute('aria-hidden', 'true');
        document.body.style.overflow = '';
    }

    function navegar(direcao) {
        var novoIndice = indiceAtual + direcao;
        if (novoIndice < 0) novoIndice = galeriaItens.length - 1;
        if (novoIndice >= galeriaItens.length) novoIndice = 0;
        abrirLightbox(novoIndice);
    }

    galeriaItens.forEach(function (item, i) {
        item.addEventListener('click', function () {
            abrirLightbox(i);
        });
    });

    if (lightbox) {
        var btnClose = lightbox.querySelector('.lightbox-close');
        var btnPrev = lightbox.querySelector('.lightbox-prev');
        var btnNext = lightbox.querySelector('.lightbox-next');

        if (btnClose) btnClose.addEventListener('click', fecharLightbox);
        if (btnPrev) btnPrev.addEventListener('click', function (e) { e.stopPropagation(); navegar(-1); });
        if (btnNext) btnNext.addEventListener('click', function (e) { e.stopPropagation(); navegar(1); });

        lightbox.addEventListener('click', function (e) {
            if (e.target === lightbox) fecharLightbox();
        });

        document.addEventListener('keydown', function (e) {
            if (!lightbox.classList.contains('active')) return;
            if (e.key === 'Escape') fecharLightbox();
            if (e.key === 'ArrowLeft') navegar(-1);
            if (e.key === 'ArrowRight') navegar(1);
        });
    }

    var observer = new IntersectionObserver(function (entries) {
        entries.forEach(function (entry) {
            if (entry.isIntersecting) {
                entry.target.classList.add('visible');
                observer.unobserve(entry.target);
            }
        });
    }, { threshold: 0.15 });

    document.querySelectorAll('.declaracao-card').forEach(function (card) {
        observer.observe(card);
    });
})();
