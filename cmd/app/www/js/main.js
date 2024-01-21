var activeProfile = null;

/**
 * Creates a profile button component
 * @param {string} name - The name of the profile
 * @returns {HTMLButtonElement} The profile button
 */
function DOMCreateProfileButton(name) {
    const button = document.createElement("button");
    button.classList.add("text-sm", "border", "shadow-sm", "font-medium", "ring-offset-background", "transition-colors", "focus-visible:outline-none", "focus-visible:ring-2", "focus-visible:ring-ring", "focus-visible:ring-offset-2", "disabled:pointer-events-none", "disabled:opacity-50", "bg-[#ffffff]", "text-black", "hover:bg-[#f4f4f5]", "rounded-md");
    button.onclick = () => {
        DOMDisableProfile(Array.from(button.parentNode.children).indexOf(button))
        DOMUpdateMods(name)
    }
    button.innerHTML = `
        <div class="p-2">
            <p>${name}</p>
        </div>
    `;
    return button;
}

/**
 * Sets the profile button at index to be disabled
 * @param {number} index - The index of the profile button
 */
function DOMDisableProfile(index) {
    const buttons = document.querySelectorAll("nav button")
    for (const button of buttons) {
        button.disabled = false
    }

    activeProfile = buttons[index].innerText
    buttons[index].disabled = true
}

/**
 * Sets the profile button to selected
 * @param {number} index - The index of the profile button
 * @returns {Promise<{buttons: HTMLButtonElement[], activeProfile: string}>} The profile button and the profile name
 */
async function DOMUpdateProfiles() {
    const profiles = await getProfiles()
    
    const buttons = []
    const nav = document.querySelector("nav")
    nav.innerHTML = ""
    nav.innerHTML += "Profiles"

    for (const profile of profiles) {
        const button = DOMCreateProfileButton(profile)
        nav.appendChild(button)
        buttons.push(button)
    }

    DOMDisableProfile(0)
    activeProfile = profiles[0]
    return {
        buttons,
        activeProfile: profiles[0]
    }
}

/**
 * Creates a mod card component
 * @param {string} modName - The name of the mod
 * @returns {HTMLDivElement} The mod card
 * @param {GetModsResponse} mod - The mod
 */
function DOMCreateModCard(mod) {
    const card = document.createElement("div")
    card.classList.add("rounded-lg", "border", "bg-card", "text-card-foreground", "shadow-sm")
    card.setAttribute("data-v0-t", "card")
    card.innerHTML = `
        <div class="p-6 flex flex-col gap-4">
            <img
            alt="Mod Image"
            class="aspect-square object-cover border border-gray-200 w-full rounded-lg overflow-hidden dark:border-gray-800"
            height="200"
            src="${mod.mod_picture}"
            width="200"
            />
            <h3 class="font-semibold">${mod.mod_name} · ${mod.mod_version}</h3>
            <p class="text-sm text-gray-500 dark:text-gray-400 break-all">
            ${mod.mod_description}
            </p>
        </div>
    `
    return card
}


/**
 * Lists all mods for profile
 * @param {string} profile - The name of the profile
 * @returns {Promise<HTMLDivElement[]>} The mod cards
 */
async function DOMUpdateMods(profile) {
    const mods = await getMods(profile)
    console.log(mods)

    const modContainer = document.getElementById("mod-list")
    modContainer.innerHTML = ""

    const DOMMods = []

    for (const mod of mods) {
        const card = DOMCreateModCard(mod)
        modContainer.appendChild(card)
        DOMMods.push(card)
    }

    return DOMMods
}

document.addEventListener("DOMContentLoaded", async function (event) {
    const { buttons, activeProfile } = await DOMUpdateProfiles();
    const mods = await DOMUpdateMods(activeProfile);
});

document.getElementById("install-new-mod").addEventListener("click", async () => {
    const modUrl = prompt("Enter the URL of the mod you want to install")
    if (!modUrl) return

    const installModResponse = await installMod(activeProfile, modUrl)
    alert(installModResponse)
    
    await DOMUpdateMods(activeProfile)
})

document.getElementById("mod-search").addEventListener("input", () => {
    const search = document.getElementById("mod-search").value.toLowerCase()
    const cards = document.querySelectorAll("[data-v0-t=card]")
    for (const card of cards) {
        if (card.innerHTML.toLowerCase().includes(search)) {
            card.style.display = "block"
        } else {
            card.style.display = "none"
        }
    }
})