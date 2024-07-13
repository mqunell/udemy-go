// TypeScript version of interfaces.go

interface Bot {
	getGreeting: () => string;
}

class EnglishBot implements Bot {
	getGreeting = () => 'hello';
}

class SpanishBot implements Bot {
	getGreeting = () => 'hola';
}

const printGreeting = (b: Bot) => {
	console.log(b.getGreeting());
};

const main = () => {
	const eb = new EnglishBot();
	const sb = new SpanishBot();

	printGreeting(eb);
	printGreeting(sb);
};

main();
