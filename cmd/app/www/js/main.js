var activeProfile = null;

function DOMCreateLoadingOverlay() {
    // Create the overlay
    const overlay = document.createElement('div');
    overlay.id = 'loading-overlay';
    overlay.classList.add('fixed', 'inset-0', 'bg-gray-600', 'bg-opacity-50', 'z-50', 'flex', 'justify-center', 'items-center');
    overlay.style.display = 'none'; // Hide by default

    // Create the loading spinner
    const loader = document.createElement('div');
    loader.classList.add('loader');
    loader.innerHTML = `
        <svg class="animate-spin -ml-1 mr-3 h-10 w-10 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 0116 0H4z"></path>
        </svg>`;

    // Append the loader to the overlay
    overlay.appendChild(loader);

    // Append the overlay to the body
    document.body.appendChild(overlay);
}

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
        document.getElementById("mod-search").value = ""
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


function DOMCreateModCard(mod) {
    const card = document.createElement("div");
    card.classList.add("rounded-lg", "border", "bg-card", "text-card-foreground", "shadow-sm", "relative", "group");
    card.setAttribute("data-v0-t", "card");

    // Create delete button as a separate element
    const deleteBtn = document.createElement("button");
    deleteBtn.classList.add("mt-auto", "inline-flex", "items-center", "justify-center", "whitespace-nowrap", "rounded-md", "text-sm", "font-medium", "transition-colors", "focus-visible:outline-none", "focus-visible:ring-2", "focus-visible:ring-ring", "focus-visible:ring-offset-2", "disabled:pointer-events-none", "disabled:opacity-50", "border", "border-input", "bg-white", "text-black", "hover:bg-gray-100", "hover:text-accent-foreground", "h-10", "px-4", "py-2");
    deleteBtn.textContent = "Delete";

    // Attach event listener to delete button
    deleteBtn.addEventListener("click", async function() {
        const confirmed = confirm(`Are you sure you want to delete the mod "${mod.mod_name}"?`);
        if (confirmed) {
            const resp = await deleteMod(activeProfile, mod.mod_path_name);
            alert(resp);
            await DOMUpdateMods(activeProfile);
        }
    });

    // Mod card content
    const content = document.createElement("div");
    content.classList.add("p-6", "flex", "flex-col", "h-full", "gap-4");

    // Image element
    const image = document.createElement("img");
    image.setAttribute("alt", "Mod Image");
    image.classList.add("aspect-square", "object-cover", "border", "border-gray-200", "w-full", "rounded-lg", "overflow-hidden", "dark:border-gray-800");
    image.setAttribute("height", "200");
    image.setAttribute("src", mod.mod_picture);
    image.setAttribute("width", "200");

    // Title element
    const title = document.createElement("h3");
    title.classList.add("font-semibold", "break-all");
    title.textContent = `${mod.mod_name} · ${mod.mod_version}`;

    // Description paragraph
    const description = document.createElement("p");
    description.classList.add("text-sm", "text-gray-500", "dark:text-gray-400", "break-all", "flex-grow", "mb-2");
    description.textContent = mod.mod_description;

    // Constructing the card content
    content.appendChild(image);
    content.appendChild(title);
    content.appendChild(description);
    content.appendChild(deleteBtn);

    // Append the content to the card
    card.appendChild(content);

    return card;
}



/**
 * Lists all mods for profile
 * @param {string} profile - The name of the profile
 * @returns {Promise<HTMLDivElement[]>} The mod cards
 */
async function DOMUpdateMods(profile) {
    const mods = await getMods(profile)
    
    // Update mods tab title
    document.getElementById("mods-tab-title").innerHTML = `Mods · ${mods.length}`

    const modContainer = document.getElementById("mod-list")
    modContainer.innerHTML = ""

    const DOMMods = []

    for (const mod of mods) {
        const card = DOMCreateModCard(mod)
        modContainer.appendChild(card)
        DOMMods.push(card)
    }

    const search = document.getElementById("mod-search").value.toLowerCase()
    DomSearchFilter(search)

    return DOMMods
}

document.addEventListener("DOMContentLoaded", async function (event) {
    const { buttons, activeProfile } = await DOMUpdateProfiles();
    await DOMUpdateMods(activeProfile);
    DOMCreateLoadingOverlay();
});

document.getElementById("install-new-mod").addEventListener("click", async () => {
    const modUrl = prompt("Enter the URL of the mod you want to install (From it's thunderstore page)");
    if (!modUrl) return;

    // Show the loading overlay
    const overlay = document.getElementById('loading-overlay');
    overlay.style.display = 'flex';

    try {
        const installModResponse = await installMod(activeProfile, modUrl);
        alert(installModResponse);
        await DOMUpdateMods(activeProfile);
    } catch (error) {
        alert('Failed to install mod: ' + error.message);
    } finally {
        // Hide the loading overlay
        overlay.style.display = 'none';
    }
});

function DomSearchFilter(query) {
    const cards = document.querySelectorAll("[data-v0-t=card]")
    for (const card of cards) {
        if (card.innerHTML.toLowerCase().includes(query)) {
            card.style.display = "block"
        } else {
            card.style.display = "none"
        }
    }
}

document.getElementById("mod-search").addEventListener("input", () => {
    const search = document.getElementById("mod-search").value.toLowerCase()
    DomSearchFilter(search)
})