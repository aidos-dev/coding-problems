<h2>
<a href="https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/">1047. Remove All Adjacent Duplicates In String</a>
</h2>

<h3>Easy</h3>

<hr>


<p>You are given a string <code>s</code> consisting of lowercase English letters. A duplicate removal consists of choosing two adjacent and equal letters and removing them.</p>

<p>We repeatedly make duplicate removals on <code>s</code> until we no longer can.</p>

<p>Return <em>the final string after all such duplicate removals have been made</em>. It can be proven that the answer is unique.</p>


<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input: s = "abbaca"</strong>
<strong>Output: "ca"</strong>
Explanation: 
For example, in "abbaca" we could remove "bb" since the letters are adjacent and equal, and this is the only possible move.  The result of this move is that the string is "aaca", of which only "aa" is possible, so the final string is "ca".
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> Input: s = "azxxzy"
<strong>Output:</strong> Output: "ay"
</pre>



<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
    <li><code>1 <= s.length <= 10<sup>5</sup></code></li>
    <li><code>s</code> consists of lowercase English letters.</li>
</ul>

