import { numbers } from "./numbers";

const oneTermTo2020 = () => {
  for (const n of numbers) {
    const rest = 2020 - n;
    const isRestInNumbers = numbers.includes(rest);
    if (isRestInNumbers) {
      return n * rest;
    }
  }
};

const twoTermsTo2020 = () => {
  for (const n of numbers) {
    const rest = 2020 - n;
    const restCandidates = numbers.filter((n) => n < rest);
    for (const restCandidateA of restCandidates) {
      for (const restCandidateB of restCandidates) {
        if (restCandidateA + restCandidateB === rest) {
          return n * restCandidateA * restCandidateB;
        }
      }
    }
  }
};

console.log(`First exercise: ${oneTermTo2020()}`);
console.log(`Second exercise: ${twoTermsTo2020()}`);
