# Booting Prompts

## System Prompt

"You are an extremely versatile and intelligent person, with advanced thinking skills and the ability to form and explain your cohesive chain of thoughts."

## **Prompt 1**: Expansion Task.

Expand the following statement into a detailed analysis with real-world applications and philosophical insights:  
"Anarchist principles challenge hierarchical structures, advocating for decentralized and cooperative systems as alternatives to authority-based governance."

Challenge: Discuss how anarchist principles apply to modern societal issues such as workplace organization, education, and community governance, while addressing potential criticisms of anarchism.

## **Prompt 2**: Compression Task.

Summarize the Four Noble Truths of Buddhism (as demonstrated below by the Buddha himself) into a clear, 100-word explanation suitable for a beginner’s guide to Buddhist philosophy.

“Mendicants, these two extremes should not be cultivated by one who has gone forth. What two? Indulgence in sensual pleasures, which is low, crude, ordinary, ignoble, and pointless. And indulgence in self-mortification, which is painful, ignoble, and pointless. Avoiding these two extremes, the Realized One understood the middle way of practice, which gives vision and knowledge, and leads to peace, direct knowledge, awakening, and extinguishment.

And what is that middle way of practice? It is simply this noble eightfold path, that is: right view, right thought, right speech, right action, right livelihood, right effort, right mindfulness, and right immersion. This is that middle way of practice, which gives vision and knowledge, and leads to peace, direct knowledge, awakening, and extinguishment.

Now this is the noble truth of suffering. Rebirth is suffering; old age is suffering; illness is suffering; death is suffering; association with the disliked is suffering; separation from the liked is suffering; not getting what you wish for is suffering. In brief, the five grasping aggregates are suffering.

Now this is the noble truth of the origin of suffering. It’s the craving that leads to future lives, mixed up with relishing and greed, taking pleasure wherever it lands. That is, craving for sensual pleasures, craving to continue existence, and craving to end existence.

Now this is the noble truth of the cessation of suffering. It’s the fading away and cessation of that very same craving with nothing left over; giving it away, letting it go, releasing it, and not clinging to it.

Now this is the noble truth of the practice that leads to the cessation of suffering. It is simply this noble eightfold path, that is: right view, right thought, right speech, right action, right livelihood, right effort, right mindfulness, and right immersion.

‘This is the noble truth of suffering.’ Such was the vision, knowledge, wisdom, realization, and light that arose in me regarding teachings not learned before from another. ‘This noble truth of suffering should be completely understood.’ Such was the vision that arose in me … ‘This noble truth of suffering has been completely understood.’ Such was the vision that arose in me …

‘This is the noble truth of the origin of suffering.’ Such was the vision that arose in me … ‘This noble truth of the origin of suffering should be given up.’ Such was the vision that arose in me … ‘This noble truth of the origin of suffering has been given up.’ Such was the vision that arose in me …

‘This is the noble truth of the cessation of suffering.’ Such was the vision that arose in me … ‘This noble truth of the cessation of suffering should be realized.’ Such was the vision that arose in me … ‘This noble truth of the cessation of suffering has been realized.’ Such was the vision that arose in me …

‘This is the noble truth of the practice that leads to the cessation of suffering.’ Such was the vision that arose in me … ‘This noble truth of the practice that leads to the cessation of suffering should be developed.’ Such was the vision that arose in me … ‘This noble truth of the practice that leads to the cessation of suffering has been developed.’ Such was the vision, knowledge, wisdom, realization, and light that arose in me regarding teachings not learned before from another.

As long as my true knowledge and vision about these four noble truths was not fully purified in these three rounds and twelve aspects, I didn’t announce my supreme perfect awakening in this world with its gods, Māras, and Divinities, this population with its ascetics and brahmins, its gods and humans.

But when my true knowledge and vision about these four noble truths was fully purified in these three rounds and twelve aspects, I announced my supreme perfect awakening in this world with its gods, Māras, and Divinities, this population with its ascetics and brahmins, its gods and humans.

Knowledge and vision arose in me: ‘My freedom is unshakable; this is my last rebirth; now there’ll be no more future lives.’”

Challenge: Retain the essence of the Four Noble Truths, including their connection to the Eightfold Path, without oversimplifying or misrepresenting key concepts.

## **Prompt 3**: Conversion Task.

Given this LLP app design, generate an SQL query to retrieve the following data from the LLP database:  
"For each bot listed in the 'Bot Manager' tab, display the bot's name, total number of prompts it has completed in the 'Ingressor' tab, and its average score for all prompts in the 'Reasoning' category. Only include bots that have completed at least five prompts in this category. Sort the results by average score in descending order."

Challenge:

- Write an optimized SQL query that uses joins across multiple tables (Bot Manager, Ingressor, Prompt Manager) while handling potential edge cases (e.g., bots with no completed prompts or missing scores).
- Provide example schemas for the tables involved and explain the query logic step-by-step, including how the query fits within the LLP app’s database structure.

```markdown
**LLP**: A lightweight LLM Benchmarking native desktop app to manage the LLMs stats and ingest outputs. (TODO)

- Tech Stack: Go, BubbleTea/Bubbles, SQLite/FTS5, Mermaid

- Tabs:

  - **(L)eaderboard** (main, or on CtrlL pressed)
    - **(I)ngressor** (on row selected and on CtrlI pressed): select a particular category, then prompt, and input the scores, speed, and output
    - **(E)gressor** (on row selected and on CtrlE pressed): view bot params and row details
    - **E(x)porter** (on CtrlX pressed): export table to json, csv, or markdown
  - **(B)ot Manager** (on CtrlB pressed): CRUD on bots, full text search, preloaded from LM Studio model list
  - **(P)rompt Manager** (on CtrlP pressed): CRUD on categories and prompts, full text search
  - **Con(d)ucer** (on CtrlD pressed): select bot, category, prompt, and then (on CtrlT pressed) will directly send request to LM studio server, and then save the output to the appropriate location
  - (save current state on CtrlS pressed and switch field via Arrows or Tab/ShiftTab, work with every tab; on CtrlY pressed at Ingressor to copy prompt)

- Directory structure:

  - `assets/`: all assets.
  - `db/`: database file, schemas, and `migration.go` to load prompt suites into db if not exist.
  - `llm_outputs/`: all LLM outputs.
  - `main.go`: glue all the tabs together.
  - `leaderboard.go`: each row is dedicated to a bot, and each column is its total points for each prompt category, another column for total points overall, and another column for average speed.
  - `ingressor.go`
  - `egressor.go`
  - `exporter.go`
  - `botman.go`
  - `promptman.go`
  - `main_test.go`: all integration tests.
  - `makefile`: all the setup and migration.
  - `conducer.go`: LM Studio OpenAI's chat endpoint: `http://127.0.0.1:1234`, `POST /v1/chat/completion`, example body:
    - `{"model": "c4ai-command-r-08-2024-i1", "stream": false, "max_tokens": -1, "messages": [{"role": "system", "content": "You are an expert translator"}, {"role": "user", "content": "Translate this text to idiomatic Vietnamese: which Sāriputta approved what the Buddha said. \r\n"}]}`
    - system prompt is a standalone table that referred to by both category and prompt.
```

## **Prompt 4**: Retrieval Task.

Identify all the philosophical assumptions in the following anarchist critique of state power and explain their significance:  
"State power inherently corrupts, as it prioritizes the maintenance of its own authority over the well-being of the people it governs."

Challenge: Uncover assumptions about human nature, authority, and societal organization, and provide an analysis of how these underpin the critique.

## **Prompt 5**: Operation Task.

Develop a set of Buddhist-inspired guidelines for ethical technology use that align with the principles of mindfulness, compassion, and non-harm.

Challenge: Address modern challenges such as social media addiction, AI ethics, and environmental sustainability, providing actionable advice rooted in Buddhist philosophy.

## **Prompt 6**: Reasoning Task.

Analyze the compatibility of anarchist philosophy with the Buddhist principle of dependent origination (paṭicca samuppāda), which emphasizes interdependence and the absence of intrinsic authority or separateness.

Challenge: Argue whether anarchism and Buddhism can be philosophically reconciled, considering their views on authority, community, and individual freedom, and provide real-world examples to support your reasoning.

## **Prompt 7**: Teaching

Can you teach me making TUI with BubbleTea, step-by-step, with clear explanations, instructions, and practice problems, so that I can use the language immediately to build a medium-sized program.

## **Prompt 8**: Programming and Explanation

Write me a program that take a positive integer at input, and output all of its prime factorization. Then explain in detail the theoretical foundations behind it and walk me through how the code work step-by-step.

## **Prompt 9**: Simple creative writing

Tell me a story about an anarchist cat, live in northen-middle India 600 bce, who peacefully standing for freedom and equality risking his own life amidst all the odds. the story should be at least 1000 words, not 1000 characters or spaces or tokens, but 1000 words.

## **Prompt 10**: Simple Translating

Can you translate for me this story into idiomatic spoken Vietnamese with proper southern dialect, semantics, and idioms?

## **Prompt 11**: Creative poetry in foreign language

Can you make a 8-line Vietnamese poem in the form of 7 words per line, called "thất ngôn bát cú đường luật".

## **Prompt 12**: Read-world translating work

### System Prompt

```text
Translate this text to idiomatic Vietnamese, which is simple and understandable to the mass, which natural spoken southern Vietnamese semantics and idioms, without many technical terms or specific Buddhist con-words. So that even a uneducated layperson can understand. Please translate it verbatim 1:1 and don't paraphrase or summarize or drop out anything.

---
```

### Prompt

Middle-length Discourse 107 - With Moggallāna the Accountant

So I have heard. At one time the Buddha was staying near Sāvatthī in the stilt longhouse of Migāra’s mother in the Eastern Monastery. Then the brahmin Moggallāna the Accountant went up to the Buddha, and exchanged greetings with him. When the greetings and polite conversation were over, he sat down to one side and said to the Buddha:

“Mister Gotama, in this stilt longhouse we can see gradual progress down to the last step of the staircase. The Buddha stops on the last step at MN 85:7.4.Among the brahmins we can see gradual progress in recitation. Memorizing the Vedic texts, a key skill of the brahmins, was so difficult that they sometimes asked the Buddha for advice (AN 5.193:8.4). Details on the gradual memorization of texts are found at Bu Pc 4:2.1.7. Texts were learned by line (teacher and student start and finish together), by going after the line (one starts, they finish together), by going after the syllable (the teacher prompts with the first syllable of the line), and by phrase (the teacher says the first phrase, the student the second).Among archers we can see gradual progress in archery. Archery (issattha) is regularly listed as a craft or livelihood (MN 14:7.3), which took skill in training (SN 56.45:1.3), and to which a mendicant is compared (AN 4.196:10.3).Among us accountants, who earn a living by accounting, we can see gradual progress in mathematics. The complexities of accounting are detailed in Kauṭilya’s Arthaśāstra 2.7. There, the “accounts” are called gāṇanikya (2.7.16).For when we get an apprentice we first make them count: ‘One one, two twos, three threes, four fours, five fives, six sixes, seven sevens, eight eights, nine nines, ten tens.’ The method of listing things up to tens is the framework of the Dasuttarasutta (DN 34). More generally, it seems to underlie the “aṅguttara principle” of organizing teachings by number.We even make them count up to a hundred. Is it possible to similarly describe a gradual training, gradual progress, and gradual practice in this teaching and training?”

“It is possible, brahmin. Suppose a deft horse trainer were to obtain a fine thoroughbred. First of all he’d make it get used to wearing the bit. See MN 65:33.1.In the same way, when the Realized One gets a person for training they first guide them like this: ‘Come, mendicant, be ethical and restrained in the monastic code, conducting yourself well and seeking alms in suitable places. Seeing danger in the slightest fault, keep the rules you’ve undertaken.’

When they have ethical conduct, the Realized One guides them further: ‘Come, mendicant, guard your sense doors. When you see a sight with your eyes, don’t get caught up in the features and details. If the faculty of sight were left unrestrained, bad unskillful qualities of covetousness and displeasure would become overwhelming. For this reason, practice restraint, protect the faculty of sight, and achieve restraint over it. When you hear a sound with your ears … When you smell an odor with your nose … When you taste a flavor with your tongue … When you feel a touch with your body … When you know an idea with your mind, don’t get caught up in the features and details. If the faculty of mind were left unrestrained, bad unskillful qualities of covetousness and displeasure would become overwhelming. For this reason, practice restraint, protect the faculty of mind, and achieve its restraint.’

When they guard their sense doors, the Realized One guides them further: ‘Come, mendicant, eat in moderation. Reflect rationally on the food that you eat: ‘Not for fun, indulgence, adornment, or decoration, but only to sustain this body, to avoid harm, and to support spiritual practice. In this way, I shall put an end to old discomfort and not give rise to new discomfort, and I will have the means to keep going, blamelessness, and a comfortable abiding.’

When they eat in moderation, the Realized One guides them further: ‘Come, mendicant, be committed to wakefulness. Practice walking and sitting meditation by day, purifying your mind from obstacles. In the first watch of the night, continue to practice walking and sitting meditation. In the middle watch, lie down in the lion’s posture—on the right side, placing one foot on top of the other—mindful and aware, and focused on the time of getting up. In the last watch, get up and continue to practice walking and sitting meditation, purifying your mind from obstacles.’

When they are committed to wakefulness, the Realized One guides them further: ‘Come, mendicant, have mindfulness and situational awareness. Act with situational awareness when going out and coming back; when looking ahead and aside; when bending and extending the limbs; when bearing the outer robe, bowl and robes; when eating, drinking, chewing, and tasting; when urinating and defecating; when walking, standing, sitting, sleeping, waking, speaking, and keeping silent.’

When they have mindfulness and situational awareness, the Realized One guides them further: ‘Come, mendicant, frequent a secluded lodging—a wilderness, the root of a tree, a hill, a ravine, a mountain cave, a charnel ground, a forest, the open air, a heap of straw.’ And they do so.

After the meal, they return from almsround, sit down cross-legged, set their body straight, and establish mindfulness in their presence. Giving up covetousness for the world, they meditate with a heart rid of covetousness, cleansing the mind of covetousness. Giving up ill will and malevolence, they meditate with a mind rid of ill will, full of sympathy for all living beings, cleansing the mind of ill will. Giving up dullness and drowsiness, they meditate with a mind rid of dullness and drowsiness, perceiving light, mindful and aware, cleansing the mind of dullness and drowsiness. Giving up restlessness and remorse, they meditate without restlessness, their mind peaceful inside, cleansing the mind of restlessness and remorse. Giving up doubt, they meditate having gone beyond doubt, not undecided about skillful qualities, cleansing the mind of doubt.

They give up these five hindrances, corruptions of the heart that weaken wisdom. Then, quite secluded from sensual pleasures, secluded from unskillful qualities, they enter and remain in the first absorption, which has the rapture and bliss born of seclusion, while placing the mind and keeping it connected. As the placing of the mind and keeping it connected are stilled, they enter and remain in the second absorption, which has the rapture and bliss born of immersion, with internal clarity and mind at one, without placing the mind and keeping it connected. And with the fading away of rapture, they enter and remain in the third absorption, where they meditate with equanimity, mindful and aware, personally experiencing the bliss of which the noble ones declare, ‘Equanimous and mindful, one meditates in bliss.’ Giving up pleasure and pain, and ending former happiness and sadness, they enter and remain in the fourth absorption, without pleasure or pain, with pure equanimity and mindfulness.

That’s how I instruct the mendicants who are trainees—who haven’t achieved their heart’s desire, but live aspiring to the supreme sanctuary from the yoke. The “trainee” (sekha) in the strict sense is restricted to those who have entered the ranks of the Noble Ones (ariyapuggala) through the realization of the four noble truths (eg. SN 48.53:3.1); that is to say, the seven Noble Ones excluding the arahant, who is an “adept” beyond training (asekha). However, in cases such as this it would be over-strict to insist that the teaching applies only to Noble Ones, as the Gradual Training is the recommended practice for all new monastics. Indeed, the burden of the text is to show how practice is taken up gradually, so above it says that this is how the Buddha “first guides them” (paṭhamaṁ evaṁ vineti, MN 107:3.3). This agrees with the Chinese parallel, which here mentions a “young monk” (MA 144 at T i 652a29). Late canonical Pali texts introduce the idea of the “ordinary person of good character” (puthujjanakalyāṇaka, eg. Cnd 8:81.2, Ps 1.1:206.2), who is said to “train” like the trainees (Pvr 1.1:3.46). The commentaries say they may be counted along with the seven Noble Ones as a “trainee” (sekkhoti puthujjanakalyāṇakena saddhiṁ satta ariyā, commentaries to Pārājika 1 and Jhānavibhaṅga).But for those mendicants who are perfected—who have ended the defilements, completed the spiritual journey, done what had to be done, laid down the burden, achieved their own goal, utterly ended the fetter of continued existence, and are rightly freed through enlightenment—these things lead to blissful meditation in the present life, and to mindfulness and awareness.”

When he had spoken, Moggallāna the Accountant said to the Buddha, “When his disciples are instructed and advised like this by Mister Gotama, do all of them achieve the ultimate goal, extinguishment, or do some of them fail?”

“Some succeed, while others fail.”

“What is the cause, Mister Gotama, what is the reason why, though extinguishment is present, the path leading to extinguishment is present, and Mister Gotama is present to encourage them, still some succeed while others fail?”

“Well then, brahmin, I’ll ask you about this in return, and you can answer as you like. What do you think, brahmin? Are you skilled in the road to Rājagaha?” Chāndogya Upaniṣad 6.14.1–2 illustrates the role of a teacher with the story of a man kidnapped, blindfolded, and abandoned in a wilderness. A kind person looses his bandage and shows him the way to Gandhāra. See too Śatapatha Brāhmaṇa 13.2.3–2, where the horse knows the way to heaven that humans do not, like one who knows the country.

“Yes, I am.”

“What do you think, brahmin? Suppose a person was to come along who wanted to go to Rājagaha. He’d approach you and say: ‘Sir, I wish to go to Rājagaha. Please point out the road to Rājagaha.’ You’d say to them: ‘Here, mister, this road goes to Rājagaha. Go along it for an hour, and you’ll see a certain village. Go along an hour further, and you’ll see a certain town. Go along an hour further and you’ll see Rājagaha with its delightful parks, woods, meadows, and lotus ponds.’ Instructed like this by you, they might still take the wrong road, heading west. Pacchāmukha means “heading west”. The present sutta is set in Sāvatthī, so for someone wanting to get to Rājagaha in the south-east, west is the wrong way. The word also appears at Thag 10.1:3.4 of the Buddha crossing the Rohiṇī river, which runs north to south; the background story says he was coming from Rājagaha, which again implies he was heading west from Koliya to Sakya.But a second person might come with the same question and receive the same instructions. Instructed by you, they might safely arrive at Rājagaha. What is the cause, brahmin, what is the reason why, though Rājagaha is present, the path leading to Rājagaha is present, and you are there to encourage them, one person takes the wrong path and heads west, while another arrives safely at Rājagaha?”

“What can I do about that, Mister Gotama? I am the one who shows the way.” The one who “shows (or explains) the way” is also discussed at Snp 1.5:3.3, along with other good and bad ascetics. See too Dhp 276:2, “the Realized Ones show the way” (akkhātāro tathāgatā).

“In the same way, though extinguishment is present, the path leading to extinguishment is present, and I am present to encourage them, still some of my disciples, instructed and advised like this, achieve the ultimate goal, extinguishment, while some of them fail. What can I do about that, brahmin? The Realized One is the one who shows the way.”

When he had spoken, Moggallāna the Accountant said to the Buddha, “Mister Gotama, there are those faithless people who went forth from the lay life to homelessness not out of faith but to earn a livelihood. They’re devious, deceitful, and sneaky. They’re restless, insolent, fickle, scurrilous, and loose-tongued. They do not guard their sense doors or eat in moderation, and they are not committed to wakefulness. They don’t care about the ascetic life, and don’t keenly respect the training. They’re indulgent and slack, leaders in backsliding, neglecting seclusion, lazy, and lacking energy. They’re unmindful, lacking situational awareness and immersion, with straying minds, witless and idiotic. Mister Gotama does not live together with these. Also at MN 5:32.1.

But there are those gentlemen who went forth from the lay life to homelessness out of faith. They’re not devious, deceitful, and sneaky. They’re not restless, insolent, fickle, scurrilous, and loose-tongued. They guard their sense doors and eat in moderation, and they are committed to wakefulness. They care about the ascetic life, and keenly respect the training. They’re not indulgent or slack, nor are they leaders in backsliding, neglecting seclusion. They’re energetic and determined. They’re mindful, with situational awareness, immersion, and unified minds; wise and clever. Mister Gotama does live together with these.

Of all kinds of fragrant root, spikenard is said to be the best. Of all kinds of fragrant heartwood, red sandalwood is said to be the best. Of all kinds of fragrant flower, jasmine is said to be the best. In the same way, Mister Gotama’s advice is the best of contemporary teachings.

Excellent, Mister Gotama! Excellent! As if he were righting the overturned, or revealing the hidden, or pointing out the path to the lost, or lighting a lamp in the dark so people with clear eyes can see what’s there, Mister Gotama has made the Teaching clear in many ways. I go for refuge to Mister Gotama, to the teaching, and to the mendicant Saṅgha. From this day forth, may Mister Gotama remember me as a lay follower who has gone for refuge for life.”
