require "csv"

def numeric?(char)
	char.match /\A[0-9]+\Z/
end

def parseSingle(content)
	note = content[0]
	mod = ""
	octave = 0
	if content.size > 1 then
		char = content[1]
		if numeric? char then
			octave = char.to_i - 1 
		else
			mod = "S"
		end
	end
	if content.size > 2 then
		ocatve = content[2].to_i - 1
	end
	"{TONE_#{note}#{mod}, #{octave}}"
end

def parse(content)
	arr = content.split(" ")
	chords = arr.collect do |item|
		parseSingle item
	end
	"\n\t&MidiSequence { []Note{ #{chords.join(", ")} } }"
end

file = File.new("./chords.csv")
csv = CSV.new(file, {:col_sep => ';', :quote_char => '"'})

csv.each do |row|
	base = parse row[0]
	major = parse row[1]
	minor = parse row[2]
	major7 = parse row[3]
	minor7 = parse row[4]
	minor7ff = parse row[5]
	dom7 = parse row[6]
	diminished = parse row[7]

	inits = [base, major, minor, major7, minor7, minor7ff, dom7, diminished].join(", ")
	puts "[]*MidiSequence{ #{inits} },"
end
