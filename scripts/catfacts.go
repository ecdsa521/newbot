package scripts

import (
	"fmt"
	"math/rand"
	"time"

	irc "github.com/thoj/go-ircevent"
)

//ScriptCatfacts ...
type ScriptCatfacts struct {
	author      string
	description string
	command     string
}

var catFacts = []string{
	"Every year, nearly four million cats are eaten in Asia",
	"On average, cats spend 2/3 of every day sleeping",
	"Unlike dogs, cats do not have a sweet tooth",
	"When a cat chases its prey, it keeps its head level",
	"The technical term for a cat's hairball is a bezoar",
	"A group of cats is called a clowder",
	"Female cats tend to be right pawed, while male cats are more often left pawed",
	"A cat cannot climb head first down a tree because its claws are curved the wrong way",
	"Cats make about 100 different sounds",
	"A cat's brain is biologically more similar to a human brain than it is to a dog's",
	"There are more than 500 million domestic cats in the world",
	"Approximately 24 cat skins can make a coat",
	"During the Middle Ages, cats were associated with witchcraft",
	"Cats are the most popular pet in North American Cats are North America's most popular pets",
	"Approximately 40,000 people are bitten by cats in the U.S.",
	"A cat's hearing is better than a dog's",
	"A cat can travel at a top speed of approximately 31 mph (49 km) over a short distance",
	"A cat can jump up to five times its own height in a single bound",
	"Some cats have survived falls of over 20 meters",
	"Researchers are unsure exactly how a cat purrs",
	"When a family cat died in ancient Egypt, family members would mourn by shaving off their eyebrows",
	"In 1888, more than 300,000 mummified cats were found an Egyptian cemetery",
	"Most cats give birth to a litter of between one and nine kittens",
	"Smuggling a cat out of ancient Egypt was punishable by death",
	"The earliest ancestor of the modern cat lived about 30 million years ago",
	"The biggest wildcat today is the Siberian Tiger",
	"The smallest wildcat today is the Black-footed cat",
	"Many Egyptians worshipped the goddess Bast, who had a woman's body and a cat's head",
	"Mohammed loved cats and reportedly his favorite cat, Muezza, was a tabby",
	"The smallest pedigreed cat is a Singapura, which can weigh just 4 lbs",
	"Cats hate the water because their fur does not insulate well when it's wet",
	"The Egyptian Mau is probably the oldest breed of cat",
	"A cat usually has about 12 whiskers on each side of its face",
	"A cat's eyesight is both better and worse than humans",
	"A cat's jaw can't move sideways, so a cat can't chew large chunks of food",
	"A cat almost never meows at another cat, mostly just humans",
	"A cat's back is extremely flexible because it has up to 53 loosely fitting vertebrae",
	"Many cat owners think their cats can read their minds",
	"Two members of the cat family are distinct from all others: the clouded leopard and the cheetah",
	"In Japan, cats are thought to have the power to turn into super spirits when they die",
	"Most cats had short hair until about 100 years ago, when it became fashionable to own cats and experiment with breeding",
	"Cats have 32 muscles that control the outer ear",
	"Cats have about 20,155 hairs per square centimeter",
	"The first cat show was organized in 1871 in London",
	"A cat has 230 bones in its body",
	"Foods that should not be given to cats include onions, garlic, green tomatoes, raw potatoes, chocolate, grapes, and raisins",
	"A cat's heart beats nearly twice as fast as a human heart",
	"Cats spend nearly 1/3 of their waking hours cleaning themselves",
	"Grown cats have 30 teeth",
	"The largest cat breed is the Ragdoll",
	"Cats are extremely sensitive to vibrations",
	"The cat who holds the record for the longest non-fatal fall is Andy",
	"The richest cat is Blackie who was left £15 million by his owner, Ben Rea",
	"Around the world, cats take a break to nap —a catnap— 425 million times a day",
	"In homes with more than one cat, it is best to have cats of the opposite sex. They tend to be better housemates.",
	"Cats are unable to detect sweetness in anything they taste",
	"Perhaps the oldest cat breed on record is the Egyptian Mau, which is also the Egyptian language's word for cat",
	"In one litter of kittens, there could be multiple father cats",
	"Teeth of cats are sharper when they're kittens. After six months, they lose their needle-sharp milk teeth",
	"Collectively, kittens yawn about 200 million time per hour",
	"According to the International Species Information Service, there are only three Marbled Cats still in existence worldwide.  One lives in the United States.",
	"Cats show affection and mark their territory by rubbing on people. Glands on their face, tail and paws release a scent to make its mark",
	"Maine Coons are the most massive breed of house cats. They can weigh up to around 24 pounds",
	"If you killed a cat in the ages of Pharaoh, you could've been put to death",
	"Most cats will eat 7 to 20 small meals a day. This interesting fact is brought to you by Nature's Recipe®",
	"Most cats don't have eyelashes",
	"Call them wide-eyes: cats are the mammals with the largest eyes",
	"Cats who eat too much tuna can become addicted, which can actually cause a Vitamin E deficiency",
	"Cats can pick up on your tone of voice, so sweet-talking to your cat has more of an impact than you think",
	"Some cats can survive falls from as high up as 65 feet or more",
	"Genetically, cats' brains are more similar to that of a human than a dog's brain",
	"If your cat's eyes are closed, it's not necessarily because it's tired. A sign of closed eyes means your cat is happy or pleased",
	"Cats CAN be lefties and righties, just like us. More than forty percent of them are, leaving some ambidextrous",
	"Cats have the skillset that makes them able to learn how to use a toilet",
	"Each side of a cat's face has about 12 whiskers",
	"Landing on all fours is something typical to cats thanks to the help of their eyes and special balance organs in their inner ear. These tools help them straighten themselves in the air and land upright on the ground.",
	"Eating grass rids a cats' system of any fur and helps with digestion",
	"Cats have 24 more bones than humans",
	"Black cats aren't an omen of ill fortune in all cultures. In the UK and Australia, spotting a black cat is good luck",
	"The Maine Coon is appropriately the official State cat of its namesake state",
	"The world's most fertile cat, whose name was Dusty, gave birth to 420 kittens in her lifetime",
	"Sometimes called the Canadian Hairless, the Sphynx is the first cat breed that has lasted this long—the breed has been around since 1966",
	"Sir Isaac Newton, among his many achievements, invented the cat flap door",
	"In North America, cats are a more popular pet than dogs. Nearly 73 million cats and 63 million dogs are kept as household pets",
	"Today, cats are living twice as long as they did just 50 years ago",
	"Outdoor cats' lifespan averages at about 3 to 5 years; indoor cats have lives that last 16 years or more",
	"Cats have the cognitive ability to sense a human's feelings and overall mood",
	"Cats prefer their food at room temperature—not too hot, not too cold",
	"Bobtails are known to have notably short tails -- about half or a third the size of the average cat",
	"A fingerprint is to a human as a nose is to a cat",
	"Cats have over 100 sounds in their vocal repertoire, while dogs only have 10",
	"Cats came to the Americas from Europe as pest controllers in the 1750s",
	"According to the Association for Pet Obesity Prevention (APOP), about 50 million of our cats are overweight",
	"Cats use their whiskers to measure openings, indicate mood and general navigation",
	"Blue-eyed cats have a high tendency to be deaf, but not all cats with blue eyes are deaf",
	"Ancient Egyptians first adored cats for their finesse in killing rodents—as far back as 4,000 years ago",
	"The color of York Chocolates becomes richer with age. Kittens are born with a lighter coat than the adults",
	"Because of widespread cat smuggling in ancient Egypt, the exportation of cats was a crime punishable by death",
	"Cats actually have dreams, just like us. They start dreaming when they reach a week old",
	"It is important to include fat in your cat's diet because they're unable to make the nutrient in their bodies on their own",
	"A cat's field of vision does not cover the area right under its nose",
	"Talk about Facetime: Cats greet one another by rubbing their noses together",
	"Cats sleep 16 hours of any given day",
	"Although it is known to be the tailless cat, the Manx can be born with a stub or a short tail",
	"A Selkirk slowly loses its naturally-born curly coat, but it grows again when the cat is around 8 months",
	"A cat's heart beats almost double the rate of a human heart, from 110 to 140 beats per minute",
	"Ragdoll cats live up to their name: they will literally go limp, with relaxed muscles, when lifted by a human",
	"Unlike most other cats, the Turkish Van breed has a water-resistant coat and enjoys being in water",
	"Webbed feet on a cat? The Peterbald's got 'em! They make it easy for the cat to get a good grip on things with skill",
	"Despite appearing like a wild cat, the Ocicat does not have an ounce of wild blood",
	"Cat's back claws aren't as sharp as the claws on their front paws",
	"A group of kittens is called a kindle, and clowder is a term that refers to a group of adult cats",
	"A third of cats' time spent awake is usually spent cleaning themselves",
	"A female cat is also known to be called a queen or a molly",
	"Want to call a hairball by its scientific name? Next time, say the word bezoar",
	"Cats have a 5 toes on their front paws and 4 on each back paw",
	"In multi-pet households, cats are able to get along especially well with dogs if they're introduced when the cat is under 6 months old and the dog is under one year old",
	"Twenty-five percent of cat owners use a blow drier on their cats after bathing",
	"Rather than nine months, cats' pregnancies last about nine weeks",
	"It has been said that the Ukrainian Levkoy has the appearance of a dog, due to the angles of its face",
	"A cat can reach up to five times its own height per jump",
	"Cats would rather starve themselves than eat something they don't like. This means they will refuse an unpalatable -- but nutritionally complete -- food for a prolonged period",
	"The Snow Leopard, a variety of the California Spangled Cat, always has blue eyes",
	"The two outer layers of a cat's hair are called, respectively, the guard hair and the awn hair",
	"When a household cat died in ancient Egypt, its owners showed their grief by shaving their eyebrows",
	"Caution during Christmas: poinsettias may be festive, but they’re poisonous to cats",
	"Most kittens are born with blue eyes, which then turn color with age",
	"A cat's meow is usually not directed at another cat, but at a human. To communicate with other cats, they will usually hiss, purr and spit.",
	"According to the Guinness World Records, the largest domestic cat litter totaled at 19 kittens, four of them stillborn",
	"As temperatures rise, so do the number of cats. Cats are known to breed in warm weather, which leads many animal advocates worried about the plight of cats under Global Warming.",
	"Cats' rough tongues enable them to clean themselves efficiently and to lick clean an animal bone",
	"Most cat litters contain four to six kittens",
	"Dogs are better than cats",
}

func init() {
	ScriptHooks["PRIVMSG"] = append(ScriptHooks["PRIVMSG"], ScriptCatfacts{
		command:     "!cat",
		author:      "moo",
		description: "Replies with random cat fact",
	}.OnMessage)

}
func randCatFact() string {
	rand.Seed(time.Now().Unix() * rand.Int63())
	return catFacts[rand.Intn(len(catFacts))]
}

//OnMessage ...
func (h ScriptCatfacts) OnMessage(event *irc.Event, irc *irc.Connection) {
	if h.command == event.Message() {
		irc.Privmsg(event.Arguments[0], randCatFact())
	}
	fmt.Printf("%v | %v\n", h.command, event.Message())
}
