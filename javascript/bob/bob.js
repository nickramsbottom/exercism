const isQuestion = message => message.endsWith('?');
const isShouting = message => message === message.toUpperCase() && (/[a-zA-Z]/.test(message));

export const hey = message => {
	const trimmed = message.trim();

	if (trimmed == '') {
		return 'Fine. Be that way!';
	}

	if (isShouting(trimmed) && isQuestion(trimmed)) {
		return 'Calm down, I know what I\'m doing!';
	}

	if (isShouting(trimmed)) {
		return 'Whoa, chill out!';
	}

	if (isQuestion(trimmed)) {
		return 'Sure.';
	}

	return 'Whatever.';
};
