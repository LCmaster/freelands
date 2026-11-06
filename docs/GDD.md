# FREELANDS
## Game Design Document

**Version:** 0.3  
**Status:** Living design document  
**Genre:** 3D Isometric Fantasy MMORPG  
**Platform:** Web Browser  
**Client Technology:** SvelteKit, TypeScript, Three.js  
**Server Technology:** Go  
**Code License:** MIT  
**Asset License:** CC0 1.0 (intended direction)

---

# 1. Game Overview

## 1.1 Concept

Freelands is a 3D isometric fantasy MMORPG set on a vast continent filled with ancient civilizations, independent kingdoms, frontier settlements and unexplored territories.

Players arrive from distant lands, bringing their own reasons for seeking a new life in Freelands.

They are not chosen heroes destined to save the world. They are adventurers who can explore, fight, trade, craft, gather resources, build relationships, discover forgotten places and establish themselves in the world.

The game emphasizes freedom of progression, meaningful exploration, player-driven economic activity and a living fantasy world.

## 1.2 Core Identity

Freelands is built around five fundamental ideas:

1. **A world worth inhabiting:** A rich fantasy setting with recognizable places, memorable inhabitants, history and mysteries.
2. **Freedom of identity:** Characters are not restricted to predefined classes. Players develop their own capabilities through their actions and choices.
3. **Meaningful activities:** Combat, exploration, crafting, gathering and trading are all legitimate ways to progress.
4. **A living multiplayer world:** Players share an environment, interact with one another and participate in its economy.
5. **Discovery over exposition:** The world reveals its history and secrets through exploration, characters, places and experiences.

## 1.3 Design Philosophy

Freelands should feel like an adventure rather than a checklist.

Players should be able to spend an evening exploring a forest, learning blacksmithing, trading between settlements, hunting dangerous creatures or helping another player.

There should be no requirement to follow a single main storyline or specialize in combat to enjoy the game.

The world exists independently of any individual player's journey.

---

# 2. Inspirations and Creative Direction

Freelands draws inspiration from several fantasy games, while developing its own identity.

| Inspiration | Design lessons |
|---|---|
| Dofus / Wakfu | Distinctive fantasy identity, mythology, memorable creatures and cultures |
| Tibia | Exploration, mystery, environmental storytelling and interconnected discoveries |
| Ragnarok Online | Memorable locations, social hubs, atmosphere and recognizable creatures |
| Skyrim | Freedom of progression, exploration and player-defined specialization |
| Torchlight II | Colourful adventure, charm and memorable environments |
| Diablo III | Isometric presentation and environmental composition |
| Path of Exile 2 | Material quality, lighting and environmental detail |

These are sources of inspiration, not templates to reproduce.

Freelands should develop original creatures, civilizations, locations, stories, visual identities and gameplay systems.

---

# 3. The World of Freelands

## 3.1 The Continent

Freelands is a vast continent surrounded by islands, archipelagos and distant seas.

It contains several different kinds of territory:

- Established kingdoms and civilizations.
- Independent cities and settlements.
- Small frontier communities.
- Ancient ruins and forgotten civilizations.
- Vast wilderness regions.
- Unmapped or poorly understood territories.
- Islands and archipelagos with their own histories.

The continent is not an empty frontier. Its inhabitants have their own cultures, histories, political relationships and understanding of the world.

Some regions are well established, while others remain difficult to access or poorly explored by outsiders.

## 3.2 Political Landscape

The continent does not belong to a single unified empire.

Its political landscape includes:

- Independent kingdoms.
- City-states.
- Duchies and principalities.
- Independent settlements.
- Unclaimed territories.
- Other forms of political organization that may emerge as the setting develops.

Relations between these powers are politically tense.

Rivalries, territorial disputes, competing interests and occasional wars exist, but most people continue living ordinary lives.

Political conflict should create opportunities for stories and gameplay without making the entire world feel like it is constantly collapsing.

## 3.3 The Outside World

The lands beyond Freelands are deliberately undefined.

People in Freelands know that other lands exist. Ships occasionally arrive from distant places, and foreign merchants, travellers, adventurers and explorers are familiar, particularly in major ports.

However, the wider world beyond the continent is outside the initial scope of the setting.

The player's homeland and the detailed geography, kingdoms and cultures beyond Freelands do not need to be established at the beginning.

This allows future expansion without requiring a complete world map or cosmology upfront.

## 3.4 The Player's Arrival

The player is a newcomer arriving from beyond the continent.

Their reason for coming is personal.

They may seek adventure, opportunity, knowledge, wealth, freedom, a new beginning or simply somewhere to go.

The opening experience introduces the player to Freelands through the arrival itself rather than a lengthy exposition sequence.

The exact starting settlement and opening location remain to be developed.

---

# 4. Character Creation and Origins

## 4.1 Playable Races

The initial playable race is **Human**.

Other races may be introduced later as the world and its civilizations develop.

The initial character system should not require a complete generalized race-management framework.

Character appearance and identity should be independent of gameplay advantages.

Humans initially have no racial bonuses, exclusive attributes, abilities or mechanical advantages.

## 4.2 Character Creation

Character creation takes place during the opening interaction with a guard.

Rather than creating a character before entering the game, the player arrives in Freelands and is questioned about their business.

The interaction introduces the player to the world while allowing them to define their character.

The player can determine:

- Gender.
- Physical appearance.
- Face.
- Hair.
- Other available visual characteristics.

The guard's dialogue can adapt to the player's selected gender and appearance.

The exact character customization options remain open.

## 4.3 Character Origins

Origins represent the character's personal reason for arriving in Freelands.

They are narrative introductions and onboarding experiences, not character classes or gameplay builds.

Potential origins include:

- Adventurer.
- Merchant.
- Scholar.
- Craftsman.
- Wanderer.
- Condemned.

These are illustrative possibilities, not a finalized list.

All origins begin at the same physical location.

### Mechanical equality

Every character starts with the same mechanical baseline:

- Same starting equipment.
- Same starting items.
- Same starting money.
- Same attributes.
- Same skills.
- Same gameplay opportunities.

Origins must not grant exclusive contacts, resources, abilities, knowledge or other mechanical advantages.

### Opening experience

The origin determines how the player answers the guard's question and introduces their first narrative experience.

Different introductions may lead to different tutorial situations or small introductory quests, but all players eventually join the same open world.

---

# 5. Character Progression

## 5.1 Classless Progression

Freelands uses a classless character progression system.

Players are not permanently assigned to predefined classes such as warrior, mage, archer or blacksmith.

Instead, characters develop through a combination of:

1. Attributes.
2. Skills.
3. Perks.
4. Knowledge.

Players gradually develop specializations through their actions and decisions.

A character can become a capable swordsman, blacksmith, alchemist, explorer or combination of several disciplines.

## 5.2 Adventurer Level

Characters earn universal experience through meaningful activities.

Potential sources include:

- Combat.
- Exploration.
- Quests.
- Gathering.
- Crafting.
- Trading.
- Discoveries.
- Other meaningful accomplishments.

Adventurer level provides universal progression opportunities, such as attribute and/or perk points.

The exact level curve and point distribution remain undecided.

## 5.3 Attributes

The initial conceptual attributes are:

| Attribute | Primary purpose |
|---|---|
| STR — Strength | Physical power, heavy weapons and carrying capacity |
| DEX — Dexterity | Accuracy, attack speed and evasion |
| INT — Intelligence | Magical power, spell complexity and alchemy |
| CON — Constitution | Health, stamina and endurance |
| WIS — Wisdom | Perception, magical resistance and knowledge-related capabilities |

Attributes represent inherent capabilities rather than learned proficiency.

Their exact formulas, secondary effects and progression values remain open.

Attributes should support different playstyles without making certain professions or builds unnecessarily dependent on a single stat.

For example, blacksmithing can benefit from physical strength and dexterity, while intelligence may support advanced compositions or techniques.

## 5.4 Skills

Skills represent what characters have learned through practice.

Examples:

**Combat**
- Swordsmanship.
- Axemanship.
- Archery.
- Magic.

**Crafting**
- Blacksmithing.
- Leatherworking.
- Tailoring.
- Alchemy.
- Woodworking.

**Gathering**
- Mining.
- Herbalism.
- Forestry.
- Hunting.
- Fishing.

Skills improve through relevant activities.

Using a sword improves swordsmanship. Crafting equipment improves blacksmithing. Gathering herbs improves herbalism.

Exact skill progression curves remain undecided.

## 5.5 Perks

Perks allow characters to specialize in how they use their capabilities.

Examples might include:

- More efficient heavy attacks.
- Improved precision with bows.
- Specialized crafting techniques.
- Better resource efficiency.
- Improved magical control.
- Specialized gathering methods.

Perks should create situational advantages rather than unconditional superiority.

Two characters with the same weapon skill can develop different combat styles.

## 5.6 Knowledge

Knowledge represents learned recipes, spells, techniques and other information.

Characters may acquire knowledge through:

- Books.
- Teachers.
- Exploration.
- Quests.
- Experimentation.
- Discoveries.
- Other characters and organizations.

Knowing how to perform an action is different from having the skill to perform it effectively.

For example, a character may learn a blacksmithing recipe but still need sufficient skill and suitable equipment to craft it successfully.

## 5.7 Anti-Grind Philosophy

Progression should reward meaningful activity rather than repetitive exploitation.

Repeated low-complexity actions may continue improving relevant skills, but their contribution to universal adventurer progression should diminish.

Meaningful accomplishments should receive greater recognition:

- Learning a new recipe.
- Discovering a resource.
- Reaching a dangerous location.
- Defeating a challenging enemy.
- Mastering a new technique.
- Completing a difficult crafting project.

Players should not be punished for enjoying professions or choosing peaceful activities.

The goal is to make progression feel natural without encouraging repetitive grinding.

---

# 6. Combat

## 6.1 Combat Philosophy

Combat is based on player decisions, equipment, skills and positioning.

There are no fixed combat classes.

Characters specialize through their choice of weapons, attributes, skills, perks and equipment.

Combat should allow different approaches while maintaining meaningful strengths, weaknesses and counterplay.

## 6.2 Primary Combat Tools

Characters can equip one active primary combat tool at a time.

Examples include:

- Swords.
- Axes.
- Bows.
- Other conventional weapons.
- Grimoires.

A character cannot actively use a conventional weapon and a grimoire simultaneously.

The active tool determines the character's available combat discipline.

## 6.3 Equipment Switching

Players can switch their active combat tool instantly, including during combat.

This provides flexibility when adapting to different enemies or situations.

Equipment weight affects combat handling and animations rather than switch duration.

For example, heavier weapons may have slower attacks and longer recovery periods.

## 6.4 Equipment Durability

Weapons and equipment can wear down through use.

Equipment may eventually break and become unusable until repaired.

This creates additional reasons to carry alternative equipment and gives repair professions an economic purpose.

Exact durability rates, repair costs and equipment damage rules remain open.

## 6.5 Combat Specialization

Combat specializations should create situational advantages rather than global power advantages.

A heavily armored fighter may be resilient but less mobile.

An archer may have range advantages but struggle when enemies close the distance.

A spellcaster may have powerful magical options but depend on suitable equipment, knowledge and resources.

These examples illustrate the intended design philosophy rather than finalized balance rules.

---

# 7. Magic

## 7.1 Magic in the World

Magic is an established part of Freelands, but it is specialized rather than ubiquitous.

People can learn magic through training and study.

Spellcasters require knowledge, practice, specialized tools and potentially magical materials.

Magic should feel like a discipline that takes effort to master rather than an ability automatically available to everyone.

## 7.2 Spellcasting

Active spellcasting requires an equipped grimoire.

The grimoire is a spellcaster's primary combat tool, functioning as an instrument through which magical knowledge and abilities are applied.

It is not necessarily the source of magical power itself.

A character carrying a magical sword or enchanted armor does not automatically gain the ability to cast spells.

## 7.3 Magical Equipment

Any character can use enchanted equipment, provided they meet its physical or equipment requirements.

Examples include:

- Flaming weapons.
- Protective armor.
- Rings that improve perception.
- Amulets with magical properties.
- Other enchanted objects.

Enchanted equipment does not require active spellcasting.

## 7.4 Magic and the World

Magic should have applications beyond combat.

Potential examples:

- Illumination.
- Revealing hidden objects.
- Activating ancient mechanisms.
- Influencing plants.
- Understanding magical creatures.
- Enchanting objects.
- Manipulating certain materials.

These possibilities should emerge naturally as the world and magic system develop.

## 7.5 Magical Affinity

Characters may possess different levels of magical affinity.

Affinity influences magical effectiveness and potential, rather than automatically granting access to a predefined class.

The exact relationship between affinity, attributes, skill and spellcasting remains undecided.

---

# 8. Equipment and Quick-Access Bar

## 8.1 Equipment

Equipment defines a character's capabilities and contributes to their identity.

It includes weapons, armor, accessories, tools and other usable items.

Equipment should provide meaningful trade-offs rather than simply replacing weaker items with universally superior versions.

## 8.2 Quick-Access Action Bar

Characters have a fixed number of quick-access slots.

Players can assign and rearrange items, but the number of slots does not increase through progression or equipment.

The bar supports:

- Consumables.
- Health and stamina potions.
- Food.
- Antidotes.
- Magical scrolls.
- Temporary enchantments.
- Other usable items.

Primary combat tools can be switched through the equipment system.

The exact number of slots and final interface remain open.

---

# 9. Quests and Narrative

## 9.1 Quest Philosophy

Quests should define required outcomes rather than forcing players to follow a single prescribed activity.

For example, a quest requesting ten wolf skins can be completed by:

- Hunting wolves.
- Purchasing skins from other players.
- Trading for the required materials.
- Obtaining the materials through other valid means.

If the player hunts wolves, they gain hunting progression.

If they acquire the materials through trade, they gain relevant trading progression.

The quest recognizes the outcome, while the underlying activity determines the progression.

## 9.2 Quest Variety

Quests should involve more than collecting items or defeating enemies.

They can include:

- Discoveries.
- Character interactions.
- Decisions.
- Investigations.
- Exploration.
- Crafting.
- Trading.
- Unexpected developments.
- Consequences.

Not every quest objective must have alternative solutions. Some discoveries, for example, should require actually finding the relevant location.

## 9.3 Quest Rewards

Potential rewards include:

- Experience.
- Skill progression.
- Money.
- Materials.
- Equipment.
- Recipes.
- Knowledge.
- Access to locations.
- Access to teachers or organizations.
- Reputation and relationships.
- New opportunities.

Quest rewards must not create permanent origin-based advantages at character creation.

## 9.4 Main Narrative

The main narrative should be an invitation, not an obligation.

Players can ignore it for extended periods and still enjoy professions, exploration, trading and other activities.

The game should support satisfying shorter sessions as well as longer adventures.

---

# 10. Professions and Economy

## 10.1 Professions

Noncombat professions are legitimate paths of progression.

Potential professions include:

- Blacksmithing.
- Leatherworking.
- Tailoring.
- Alchemy.
- Woodworking.
- Mining.
- Herbalism.
- Forestry.
- Hunting.
- Fishing.
- Trading.

Professions should have meaningful relationships with one another.

For example:

Hunter → Leatherworker → Merchant → Adventurer.

A player who gathers materials can supply a crafter, who produces equipment that another player purchases.

This creates opportunities for player-driven economic activity.

## 10.2 Player Economy

The economy should encourage interaction and specialization.

Players should be able to acquire goods through their own activities or through other players.

Trading should be a meaningful way to progress rather than simply a way to transfer money.

## 10.3 Economic Philosophy

The initial design direction emphasizes:

- Useful demand for crafted and gathered materials.
- Meaningful reasons for players to trade.
- Opportunities for specialized producers.
- Accessibility for new players.
- Multiple ways to acquire necessary quest and crafting materials.

Detailed market structures, NPC merchants, resource availability and economic balancing remain to be developed.

---

# 11. Multiplayer and Social Systems

## 11.1 Shared World

Freelands is a multiplayer game from its first playable version.

Players share the same world and can encounter one another naturally.

The server is authoritative and responsible for validating gameplay actions and maintaining shared state.

## 11.2 Player Interaction

The world should encourage players to interact through:

- Trading.
- Cooperation.
- Shared exploration.
- Combat.
- Crafting.
- Social activities.

The exact interaction systems will be introduced as the multiplayer foundation develops.

## 11.3 Future Social Features

Potential future systems include:

- Friends.
- Parties.
- Guilds.
- Public and private chat.
- Player trading.
- Cooperative activities.
- Organizations and factions.

These systems are not required for the initial playable prototype.

---

# 12. Housing and Player Property

## 12.1 Player-Owned Housing

Players can eventually purchase their own houses.

Housing provides a personal place in the world and creates another form of long-term progression.

Potential benefits include:

- Secure storage.
- Workshop and crafting facilities.
- Furniture and customization.
- Trophy displays.
- Social spaces.
- A personal residence.

## 12.2 Economic Responsibilities

Owning property has financial responsibilities.

Potential expenses include:

- Purchase price.
- Property taxes.
- Maintenance.
- Optional utilities or services.

Property ownership should be a meaningful choice, not a mandatory progression system.

Players who prefer travelling can use inns or rented rooms without owning property.

## 12.3 Housing Design Principles

Housing should provide meaningful benefits without becoming an unavoidable financial burden.

The exact tax system, ownership rules, storage limits and inactivity policies remain open.

---

# 13. Exploration and World Activities

Exploration is one of the central activities of Freelands.

Players should have reasons to visit places beyond their immediate settlement.

Potential exploration activities include:

- Discovering hidden locations.
- Exploring ruins.
- Finding rare resources.
- Encountering unfamiliar creatures.
- Investigating historical landmarks.
- Discovering new settlements.
- Finding forgotten routes.
- Uncovering local stories and mysteries.

Not every location needs a quest.

Some places should exist simply because they belong in the world.

Discovery should be rewarding in itself.

---

# 14. Creatures and the Natural World

Creatures should have identities beyond their combat statistics.

Their behaviour, habitats, ecology and relationships with the environment should help make them memorable.

Potential creature characteristics include:

- Preferred habitats.
- Behaviour patterns.
- Social structures.
- Defensive responses.
- Hunting behaviour.
- Interactions with other species.
- Relationships with local inhabitants.

The world should not feel like a collection of enemies placed randomly for players to defeat.

Creature design should connect naturally with geography, ecology, culture and regional identity.

Specific species, classifications and creature mechanics remain open.

---

# 15. Visual and Presentation Direction

## 15.1 Camera and Perspective

Freelands uses a 3D isometric presentation.

The camera should provide a clear view of the environment while maintaining a strong sense of depth.

Camera rotation and zoom are potential features to be explored during development.

The final camera projection and controls remain subject to testing.

## 15.2 Art Direction

The visual direction is colourful stylized realism.

The intended appearance combines:

- Believable character proportions.
- Stylized but detailed materials.
- Rich medieval fantasy architecture.
- Lush vegetation.
- Distinctive equipment.
- Carefully composed environments.
- Atmospheric lighting.
- Controlled, vibrant colours.

The game should feel inviting and adventurous without becoming excessively cartoonish.

## 15.3 Environment Design

The world is primarily handcrafted.

Towns, buildings, landmarks, interiors, dungeons, forests, props and points of interest should be deliberately designed and assembled.

Procedural techniques may support placement, variation and environmental detail, but the world itself is not intended to be procedurally generated.

Each region should have a recognizable identity.

---

# 16. Interface and User Experience

The interface should support both desktop and mobile browser experiences.

The initial development focus is desktop, with mobile compatibility considered during interface architecture.

The UI should reinforce the fantasy-adventure identity rather than feeling like a generic web application.

Potential interface elements include:

- Character status.
- Health and stamina.
- Inventory.
- Equipment.
- Quick-access action bar.
- Character progression.
- Skills and professions.
- Quest journal.
- World map.
- Trading interface.
- Social interface.

The exact HUD layout, visual language and interaction patterns remain to be designed.

---

# 17. Technical Direction

## 17.1 Architecture

The project uses a client-server architecture.

**Client**
- SvelteKit.
- TypeScript.
- Three.js.

**Server**
- Go.
- Authoritative game simulation.
- WebSocket communication.

The client is responsible for presentation, input and rendering.

The server is responsible for authoritative gameplay state and validation.

## 17.2 Multiplayer Authority

Clients send intentions and requests.

The server validates those requests and determines their results.

Modified clients must not be able to determine authoritative player positions, item ownership, combat outcomes or economic transactions.

## 17.3 Initial Development Scope

The first multiplayer prototype focuses on:

- Server startup.
- Player connections.
- Shared world state.
- Player spawning.
- Authoritative movement.
- State synchronization.
- Player disconnection.

Persistent storage, authentication and gameplay systems are introduced progressively.

---

# 18. Open Source and Licensing

Freelands is intended to be an open-source project.

The project aims to make the following available:

- Client source code.
- Server source code.
- Game logic.
- Development tools.
- Original game assets.

## 18.1 Code License

The selected source code license is MIT.

## 18.2 Asset License

Original game assets are intended to use CC0 1.0.

This direction is intended to allow broad reuse and modification.

Third-party assets and contributions must have their own rights and licenses checked before inclusion.

## 18.3 Official and Community Servers

The project will maintain official servers under project control.

Community-hosted modified servers should also be permitted and encouraged.

Security must rely on server-side authority and validation, not secrecy of the client implementation.

---

# 19. Development Principles

The following principles guide future design decisions.

1. **Freedom over predefined classes:** Characters develop through their actions and decisions.
2. **Multiple valid playstyles:** Combat, crafting, gathering, trading and exploration should all matter.
3. **Meaningful progression:** Reward learning, difficulty, discovery and accomplishment.
4. **A coherent world:** Cultures, creatures, geography, history and gameplay should influence one another.
5. **Discovery over exposition:** Let players uncover the world rather than explaining everything upfront.
6. **No chosen-one requirement:** Players are adventurers, not predetermined saviours of the universe.
7. **Player-driven interaction:** Encourage meaningful cooperation and economic relationships.
8. **Accessible onboarding:** Introduce the world and its systems gradually.
9. **Expandable design:** New regions, races, systems and stories should be introduced without requiring the entire world to be defined upfront.
10. **Avoid unnecessary complexity:** Build systems when they solve real gameplay or technical needs.

---

# 20. Initial Development Scope

The initial playable version should not attempt to implement the entire game.

The first objective is to establish a functional multiplayer environment.

## Initial prototype

- Human playable characters.
- Basic character representation.
- Shared multiplayer world.
- Player spawning.
- Authoritative movement.
- Player visibility and synchronization.
- Basic environment and ground.
- Initial camera and controls.

## Subsequent development

After the multiplayer foundation is stable, development can expand into:

- Character creation and persistence.
- Basic combat.
- Equipment and inventory.
- Attributes and skills.
- Gathering and crafting.
- NPCs and interactions.
- Quests.
- Trading.
- Exploration and world content.
- Housing and social systems.

The order and scope of these milestones will evolve as the prototype is tested.

---

# 21. Open Design Questions

These topics are intentionally not finalized.

| Area | Question |
|---|---|
| World | What is the first starting settlement? |
| Geography | What are the first regions surrounding it? |
| Races | Which additional playable races will eventually exist? |
| Character creation | What appearance customization options are available? |
| Progression | How exactly are attribute and perk points awarded? |
| Skills | What are the final skill categories and progression curves? |
| Combat | What are the detailed attack, defence and damage mechanics? |
| Magic | How do affinity, spell knowledge and magical materials interact? |
| Equipment | What are the final equipment slots and durability rules? |
| Quests | How are quest chains and long-term narrative consequences structured? |
| Economy | How do markets, NPC merchants and player trading interact? |
| Housing | How do ownership, taxation and storage work? |
| Multiplayer | How are parties, guilds and world regions organized? |
| Creatures | What are the first regional ecosystems and creature families? |
| Camera | What are the final camera controls and projection? |
| Interface | What is the final HUD and visual design system? |

These questions should be answered when they become relevant to worldbuilding, gameplay design or implementation.

They are not prerequisites for developing the initial multiplayer prototype.

---

# 22. Project Vision

Freelands is a fantasy world players can make their own.

Some will become skilled warriors. Others will master magic, establish businesses, craft equipment, explore forgotten ruins or simply travel between settlements.

Their journeys will be different, even though they share the same world.

The goal is not to create a game where everyone follows the same path.

It is to create a world worth returning to, where players can discover their own place in it.