package main

// Find the Greatest Common Divisor (GCD) of two numbers using Euclid’s algorithm.
// Time complexity: O(log(min(a, b)) — super efficient

// The Euclidean Algorithm says:

//     💡 GCD(a, b) = GCD(b, a % b)
//     Keep repeating this until b becomes 0. At that point, a is the GCD.

// Intuition
// Imagine:
//     You have a = 48, b = 18
//     You know 18 divides 48 with remainder 12 → because 48 = 2×18 + 12
//     Now, ANY number that divides both 48 and 18 will also divide 12.
//     So we switch: GCD(48, 18) = GCD(18, 12)
//     Repeat the process...
//   It’s basically like:
//     "Let’s shrink the problem to a smaller and smaller one until it's dead simple."

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b,a%b)
}

// Formula: LCM(a, b) = (a * b) / GCD(a, b)
func lcm(a,b int) int {
	return (a*b)/gcd(a,b)
}

// Two numbers are co-prime if their GCD is 1.
func isCoPrime(a, b int) bool {
	return gcd(a, b) == 1
}

// func main() {
// 	fmt.Println(gcd(48,18))
// }