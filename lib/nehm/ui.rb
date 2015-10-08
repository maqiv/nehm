module Nehm
  module UI

    # TODO: add Highline features to UI module

    def self.error(msg)
      puts "#{msg}\n".red
    end

    def self.newline
      puts "\n"
    end

    def self.say(msg)
      puts msg
    end

    def self.success(msg)
      puts msg.green
    end

    def self.term(msg = nil)
      if msg
        abort msg.red
      else
        exit
      end
    end

    def self.warning(msg)
      puts "#{msg}".yellow
    end

  end
end