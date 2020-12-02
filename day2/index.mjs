import { passwords } from "./passwords.mjs";

const countValidPasswords = (passwords) => (policy) =>
  passwords.filter((p) => p && policy(p)).length;

const byPositionPolicy = (passwordPolicy) => {
  const {
    min: firstPosition,
    max: secondPosition,
    letter,
    password
  } = getPasswordAndPolicy(passwordPolicy);
  const passwordLetters = password.split("");
  const inFirst = passwordLetters[firstPosition - 1] === letter;
  const inSecond = passwordLetters[secondPosition - 1] === letter;
  return !inFirst === inSecond;
};

const inRangePolicy = (passwordPolicy) => {
  const { min, max, letter, password } = getPasswordAndPolicy(passwordPolicy);
  const indexByLetter = countLettersOccurences(password);
  const passwordLetterOcc = indexByLetter[letter];
  if (!passwordLetterOcc) return false;
  return passwordLetterOcc >= min && passwordLetterOcc <= max;
};

const getPasswordAndPolicy = (passwordPolicy) => {
  const policyExtractRegex = /(\d+)-(\d+)\s(\w):\s(\w+)/;
  const [, min, max, letter, password] = passwordPolicy.match(
    policyExtractRegex
  );
  return {
    min: parseInt(min, 10),
    max: parseInt(max, 10),
    letter,
    password
  };
};

const countLettersOccurences = (letters) =>
  letters
    .split("")
    .reduce(
      (acc, l) => (acc[l] ? { ...acc, [l]: acc[l] + 1 } : { ...acc, [l]: 1 }),
      {}
    );

const shoeKeeperPassword = countValidPasswords(passwords);

console.log(`First exercise: ${shoeKeeperPassword(inRangePolicy)}`);
console.log(`Second exercise: ${shoeKeeperPassword(byPositionPolicy)}`);
