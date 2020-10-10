const isQuestion = message => message.trim().endsWith('?');
const isShouting = message => message == message.toUpperCase() && (/[A-Z]/.test(message));
const isSilence = message => message.trim() == '';

export const hey = message => {
	if (isSilence(message)) {
		return 'Fine. Be that way!';
	}

	if (isShouting(message) && isQuestion(message)) {
		return 'Calm down, I know what I\'m doing!';
	}

	if (isShouting(message)) {
		return 'Whoa, chill out!';
	}

	if (isQuestion(message)) {
		return 'Sure.';
	}

	return 'Whatever.';
};
