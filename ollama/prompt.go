package ollama

import "fmt"

const GeneralHeader = "## system\nYou are a creative and helpful %s programming assistant. I will provide you with some information about a variable, function, method, struct, class, or interface, and you will suggest a concise and informative name for it, considering the specific conventions and best practices of %s. Please provide a code block with the name you suggest."

const (
	LanguageGo         = "Go"
	LanguagePython     = "Python"
	LanguageTypescript = "Typescript"
)

var languageList = []string{LanguageGo, LanguagePython, LanguageTypescript}

func GetLanguageList() []string {
	clone := make([]string, len(languageList))
	copy(clone, languageList)
	return clone
}

const GolangRule = `## rules
1. **Use Camel Case:** Go uses Camel Case for naming functions and methods. This means that each word in the name starts with a capital letter and the words are concatenated without spaces. For example, 'calculateTotalPrice' or 'getUserData'.
2. **Visibility through Capitalization:** The first letter of a function or method name determines its visibility:
    - **Uppercase first letter:** Exported (accessible from other packages).
    - **Lowercase first letter:** Unexported (accessible only within the same package).
3. **Use Meaningful Names:** Choose names that clearly indicate the function's or method's purpose. For instance, a function that calculates the age of a user could be named 'calculateAge' or 'getAge'.
4. **Use Verbs or Verb Phrases:** Functions perform actions, so their names should typically include a verb or verb phrase. Examples: 'readFile', 'writeData', 'connectToServer'.
5. **Keep Names Concise and Clear:** Strive for names that are concise and easy to understand. Overly long or complex names can hinder readability.
6. **Avoid Abbreviations:** Use complete words instead of abbreviations whenever possible. Abbreviations can be confusing and reduce clarity.
7. **Maintain Consistency:** Adhere to a consistent naming convention throughout your project. This improves code readability and maintainability.
8. **Context as the First Argument:** If a function or method uses a 'context.Context', it should be the first argument. This allows for consistent context propagation and cancellation throughout your code.
9. **Error as the Last Return Type:** In Go, it's conventional to return an 'error' type as the last return value. This allows for easy error checking and handling.
10. **Boolean-Returning Methods:** When naming methods that return a boolean value, consider using prefixes like 'Is', 'Exists', 'Remain', 'Can', 'May', 'Any', or 'All' to clearly indicate the method's purpose and the expected return type.
11. **Avoid Package Names in Names:** Function and method names should not include the package name. This is because the package name is already specified when calling the function or method.`

func MakeRequest(input string) string {
	return fmt.Sprintf("## Requirement\n%s", input)
}

func MakeGoPrompt(request string) string {
	return fmt.Sprintf(GeneralHeader, "Go", "Go") + "\n\n" + GolangRule + "\n\n" + request
}
