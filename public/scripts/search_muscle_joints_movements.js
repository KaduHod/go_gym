function removeAcentos(str) {
    return str.normalize("NFD").replace(/[\u0300-\u036f]/g, "");
}

const value_includes_target = (needle, target) => {
    needle = removeAcentos(needle).toLowerCase();
    if (Array.isArray(target)) {
        const filtered_targets = target.filter(term => {
            return needle.includes(removeAcentos(term.trim().normalize()).toLowerCase());
        });
        return filtered_targets.length > 0;
    }
    return needle.includes(removeAcentos(target).toLowerCase());
}

const searchGroup = () => document.getElementById("search-group");
const searchJoint = () => document.getElementById("search-joint");
const searchPortion = () => document.getElementById("search-portion");
const searchMovement = () => document.getElementById("search-movement");
const muscleGroups = () => document.querySelectorAll(".muscle-group");

function filterMuscles() {
    const queryGroup = removeAcentos(searchGroup().value.toLowerCase());
    const queryJoint = removeAcentos(searchJoint().value.toLowerCase());
    const queryPortion = removeAcentos(searchPortion().value.toLowerCase());
    const queryMovement = removeAcentos(searchMovement().value.toLowerCase());

    let foundResults = false;
    console.log("filtros", {queryGroup, queryJoint, queryPortion, queryMovement})
    // Se todos os filtros estiverem vazios, exiba todos os cards
    if (queryGroup.length === 0 && queryJoint.length === 0 && queryPortion.length === 0 && queryMovement.length === 0) {
        muscleGroups().forEach(group => {
            group.style.display = "block";
            group.querySelectorAll(".muscle-portion").forEach(portion => {
                portion.style.display = "block";
                portion.querySelectorAll(".muscle-card").forEach(card => {
                    card.style.display = "block";
                });
            });
        });
        return;
    }

    muscleGroups().forEach(group => {
        const musclePortions = group.querySelectorAll(".muscle-portion");
        const muscleGroupName = removeAcentos(group.dataset.muscleGroup.toLowerCase().normalize());
        let hasVisiblePortions = false;

        musclePortions.forEach(portion => {
            const musclePortionName = removeAcentos(portion.dataset.musclePortion.toLowerCase().normalize());
            const cards = portion.querySelectorAll(".muscle-card");
            let hasVisibleCards = false;

            cards.forEach(card => {
                const jointName = removeAcentos(card.dataset.joint.toLowerCase().normalize());
                const movementName = removeAcentos(card.dataset.movement.toLowerCase().normalize());

                // Verifica se o card deve ser exibido com base nos filtros aplicados
                if (
                    (queryGroup.length === 0 || value_includes_target(muscleGroupName, queryGroup)) &&
                    (queryPortion.length === 0 || value_includes_target(musclePortionName, queryPortion)) &&
                    (queryJoint.length === 0 || value_includes_target(jointName, queryJoint)) &&
                    (queryMovement.length === 0 || value_includes_target(movementName, queryMovement))
                ) {
                    card.style.display = "block";
                    hasVisibleCards = true;
                    if (!foundResults) foundResults = true;
                    if (!hasVisiblePortions) hasVisiblePortions = true;
                } else {
                    card.style.display = "none";
                }
            });

            portion.style.display = hasVisibleCards ? "block" : "none";
        });

        // Exibe o grupo se tiver porções visíveis
        group.style.display = hasVisiblePortions ? "block" : "none";
    });
}

document.addEventListener("DOMContentLoaded", function () {
    const elements = document.querySelectorAll('.fade-in');
    const filterInputs = document.querySelectorAll('.filter-input');

    const observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                // Quando o elemento entra na tela, adicionamos a classe 'fade-in' para a animação
                entry.target.classList.add('fade-in');
            }
        });
    }, {
        threshold: 0.5 // O elemento será observado quando estiver 50% visível
    });

    elements.forEach((element) => {
        observer.observe(element);
    });

    // Função para reexecutar a animação quando o filtro for aplicado
    filterInputs.forEach(input => {
        input.addEventListener('input', function () {
            // Encontrar os cards visíveis após o filtro
            const visibleCards = document.querySelectorAll('.muscle-card');

            visibleCards.forEach(card => {
                if (card.offsetParent !== null) { // Verifica se o card está visível
                    // Reaplicar a animação removendo e adicionando a classe fade-in
                    card.classList.remove('fade-in');
                    // Forçar reflow para reiniciar a animação
                    void card.offsetWidth;  // Esse comando força um reflow e reinicia a animação
                    card.classList.add('fade-in');
                }
            });
        });
    });
});

