# 


<div class="brinza-task-description">
<p>Strugacarro is a planet whose year is divided into four seasons: winter, spring, summer and autumn in that order. A year has N days and every season lasts for exactly N/4 days. The year starts on the first day of winter and ends on the last day of autumn.</p>
<p>Given the history of temperatures from the previous year, find the season with the highest amplitude of temperatures. The amplitude is the difference between the highest and lowest temperatures over the given period. Assume that all seasons within one year have different temperature amplitudes.</p>
<p>Write a function:</p>

<blockquote><p style="font-family: monospace; font-size: 9pt; display: block; white-space: pre-wrap"><tt>func Solution(T []int) string</tt></p></blockquote>
<p>that, given an array T of N integers denoting the temperatures on all days of the year, returns a string with the name of the season with the highest temperature amplitude (one of the following: "<tt style="white-space:pre-wrap">WINTER</tt>", "<tt style="white-space:pre-wrap">SPRING</tt>", "<tt style="white-space:pre-wrap">SUMMER</tt>", "<tt style="white-space:pre-wrap">AUTUMN</tt>").</p>

<p>For example, given T = [−3, −14, −5, 7, 8, 42, 8, 3]:</p>
<p><img class="inline-description-image" src="https://codility-frontend-prod.s3.amazonaws.com/media/task_static/four_seasons/static/images/auto/91f740984ca835253b8616ac6e646f98.png"></p>
<p>the function should return "<tt style="white-space:pre-wrap">SUMMER</tt>", since the highest amplitude (34) occurs in summer.</p>
<p>Given T = [2, −3, 3, 1, 10, 8, 2, 5, 13, −5, 3, −18]:</p>
<p><img class="inline-description-image" src="https://codility-frontend-prod.s3.amazonaws.com/media/task_static/four_seasons/static/images/auto/c529a8402c9a97065c7217296d3dd69a.png"></p>
<p>the correct answer is "<tt style="white-space:pre-wrap">AUTUMN</tt>" (amplitude equals 21).</p>
<p>Assume that:</p>
<blockquote><ul style="margin: 10px;padding: 0px;"><li>The number of elements in the array is divisible by 4;</li>
<li>each element of array T is an integer within the range [<span class="number">−1,000</span>..<span class="number">1,000</span>];</li>
<li>N is an integer within the range [<span class="number">8</span>..<span class="number">200</span>];</li>
<li>Amplitudes of all seasons are distinct.</li>
</ul>
</blockquote><p>In your solution, focus on <b><b>correctness</b></b>. The performance of your solution will not be the focus of the assessment.</p>
</div>


# Solution
WINTER
SPRING
SUMMER
AUTUMN

Case 1, T = [−3, −14, −5, 7, 8, 42, 8, 3] HA=34 return  SUMMER 

Case 2, T = [2, −3, 3, 1, 10, 8, 2, 5, 13, −5, 3, −18] HA=21 AUTUMN


Get how many numbers will be in a season.
Divide T lenght by 4.

In case 2, lenT =12, so 3 temparatures.

First loop to go through all the temperatures.
for i =0; i< lenT; i+temperatures

Second loop to find the highest and the lowest numbers.

h=0 , l=0
for j = i; i < temperatures; i++ {

We cannot take 0 as there are -ve temperature,

So we assign the first numbers as h and l

 






